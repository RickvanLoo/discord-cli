package state

import "github.com/Rivalo/discordgo_cli"

//State is the current state of the attached client
type State struct {
	Username string
	Password string
	Guild    *discordgo.Guild
	Channel  *discordgo.Channel
	Users    []*discordgo.User
	Messages []*discordgo.Message
}
