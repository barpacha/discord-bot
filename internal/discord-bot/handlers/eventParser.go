package handlers

import "github.com/bwmarrin/discordgo"

func (r *Resolver) EventParser(session *discordgo.Session, message *any) {
	*message = (*message).(*discordgo.Event).Struct
}
