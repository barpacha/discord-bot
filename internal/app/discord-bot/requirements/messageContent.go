package requirements

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
)

func MessageContent(templateString string) func(*discordgo.Session, any) bool {
	return func(session *discordgo.Session, message any) bool {
		m := message.(*discordgo.MessageCreate)
		return len(regexp.MustCompile(templateString).FindString(m.Content)) == len(m.Content)
	}
}
