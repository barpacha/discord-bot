package discord_bot

import (
	"context"
	"database/sql"
	"discord-bot/config"
	"discord-bot/internal/app/discord-bot/handlers"
	"discord-bot/internal/app/discord-bot/store"
	"discord-bot/internal/pkg/discord"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func Run(config config.Main) error {
	db, err := sql.Open("postgres", config.Store.Database)
	if err != nil {
		log.Fatalf("store error: %s", err)
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(config.Store.ConnMaxLifetimeMinutes))
	db.SetMaxOpenConns(config.Store.MaxOpenConns)
	db.SetMaxIdleConns(config.Store.MaxIdleConns)

	if err := db.Ping(); err != nil {
		log.Fatalf("store error: %s", err)
	}
	resolver := &handlers.Resolver{Store: store.New(db)}
	client, err := discord.NewClient(config.Discord.Token, Router(resolver))
	if err != nil {
		return fmt.Errorf("new discord client error: %w", err)
	}
	ctx := gracefulContext()
	err = client.Start(ctx)
	if err != nil {
		return fmt.Errorf("new discord client start error: %w", err)
	}
	return nil
}

func gracefulContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	irqSig := make(chan os.Signal, 1)
	go func() {
		<-irqSig
		cancel()
	}()
	return ctx
}
