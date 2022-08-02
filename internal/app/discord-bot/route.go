package discord_bot

import (
	"discord-bot/internal/app/discord-bot/handlers"
	"discord-bot/internal/app/discord-bot/requirements"
	"discord-bot/internal/pkg/discord"
	"github.com/bwmarrin/discordgo"
)

func Router(resolver *handlers.Resolver) *discord.Router {
	return discord.NewRouter().AddBranch(
		requirements.Type(&discordgo.Event{}),
		discord.NewRouter().AddAction(resolver.EventParser).
			AddBranch(
				requirements.Type(&discordgo.MessageCreate{}),
				discord.NewRouter().
					AddBranch(
						requirements.NotMe,
						discord.NewRouter().
							AddBranch(
								requirements.MessageContent("Ping"),
								discord.NewRouter().AddAction(resolver.Ping),
							).
							AddBranch(
								requirements.MessageContent("Pong"),
								discord.NewRouter().AddAction(resolver.Pong),
							),
					),
			),
	)
}
