package discord

import "github.com/bwmarrin/discordgo"

type Handler func(*discordgo.Session, *any)

type HandlerMessageCreate func(*discordgo.Session, discordgo.MessageCreate)
