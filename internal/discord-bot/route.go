package discord_bot

import (
	"discord-bot/internal/discord-bot/handlers"
	"discord-bot/pkg/discord"
	"github.com/bwmarrin/discordgo"
)

func Router(resolver *handlers.Resolver) *discord.Router {
	return discord.NewRouter().
		AddBranch(discord.NewBranch(
			func(session *discordgo.Session, a any) bool {
				return false
			},
			discord.NewRouter().AddAction(resolver.PingPong),
		)).
		AddBranch(discord.NewBranch(
			func(session *discordgo.Session, a any) bool {
				return true
			},
			discord.NewRouter(),
		))
}
