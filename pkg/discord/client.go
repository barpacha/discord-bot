package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
)

type Client struct {
	session *discordgo.Session
	router  *Router
}

func NewClient(token string, router *Router) (*Client, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Client{
		session: session,
		router:  router,
	}, nil
}

func (c *Client) Start(ctx context.Context) error {
	errChan := make(chan error)
	c.session.AddHandler(func(session *discordgo.Session, message any) {
		c.router.Serve(session, message)
	})
	go func() {
		errChan <- c.session.Open()
	}()
	<-ctx.Done()
	c.Close()
	return <-errChan
}

func (c *Client) RouterCreateMessage() *Router {
	return c.router
}

func (c *Client) Close() {
	_ = c.session.Close()
}
