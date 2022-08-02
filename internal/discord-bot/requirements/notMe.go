package requirements

import (
	"github.com/bwmarrin/discordgo"
)

func NotMe(session *discordgo.Session, message any) bool {
	m := message.(*discordgo.MessageCreate)
	if m.Author.ID == session.State.User.ID {
		return false
	}
	return true
}
