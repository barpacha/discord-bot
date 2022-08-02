package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (r *Resolver) Ping(session *discordgo.Session, message *any) {
	m := (*message).(*discordgo.MessageCreate)
	_, err := session.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Printf("send message in channel %s", m.ChannelID)
	}
}

func (r *Resolver) Pong(session *discordgo.Session, message *any) {
	m := (*message).(*discordgo.MessageCreate)
	_, err := session.ChannelMessageSend(m.ChannelID, "Ping!")
	if err != nil {
		log.Printf("send message in channel %s", m.ChannelID)
	}
}
