package discord

import "github.com/bwmarrin/discordgo"

type BranchRequirement func(*discordgo.Session, any) bool
