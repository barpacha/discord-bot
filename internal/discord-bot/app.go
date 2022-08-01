package discord_bot

import (
	"context"
	"discord-bot/config"
	"discord-bot/internal/discord-bot/handlers"
	"discord-bot/pkg/discord"
	"fmt"
	"os"
)

func Run(config config.Main) error {
	resolver := &handlers.Resolver{}
	client, err := discord.NewClient(config.Discord.Token, Router(resolver))
	if err != nil {
		return fmt.Errorf("new discord client error: %w", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	irqSig := make(chan os.Signal, 1)
	go func() {
		<-irqSig
		cancel()
	}()
	err = client.Start(ctx)
	if err != nil {
		return fmt.Errorf("new discord client start error: %w", err)
	}
	return nil
}
