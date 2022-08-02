package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (r *Resolver) PingPong(session *discordgo.Session, message *any) {
	m := (*message).(*discordgo.MessageCreate)
	_, err := session.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Printf("send message in channel %s", m.ChannelID)
	}
}
