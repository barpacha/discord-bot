package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (r *Resolver) PingPong(session *discordgo.Session, message any) {
	m := message.(discordgo.MessageCreate)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == session.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		session.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		session.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
