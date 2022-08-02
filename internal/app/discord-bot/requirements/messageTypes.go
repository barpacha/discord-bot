package requirements

import (
	"github.com/bwmarrin/discordgo"
	"reflect"
)

func Type(messageType any) func(*discordgo.Session, any) bool {
	return func(session *discordgo.Session, message any) bool {
		return reflect.TypeOf(messageType) == reflect.TypeOf(message)
	}
}
