package requirements

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func (r *Resolver) serverActiveMessageCreate(session *discordgo.Session, message any) bool {
	m := message.(*discordgo.MessageCreate)
	serverId, err := strconv.Atoi(m.GuildID)
	if err != nil {
		return false
	}
	server, err := r.Store.Server().Get(context.Background(), uint(serverId))
	if err != nil {
		return false
	}
	if server == nil {
		return false
	}
	return true
}
