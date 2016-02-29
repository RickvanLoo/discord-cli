package DiscordState

import "github.com/Rivalo/discordgo_cli"

//State is the current state of the attached client
type State struct {
	Guild    *discordgo.Guild
	Channel  *discordgo.Channel
	Channels map[string]*discordgo.Channel
	Members  map[string]*discordgo.Member
	Messages map[string]*discordgo.Message
	Session  *Session
}

//Session contains the 'state' of the attached server
type Session struct {
	Username  string
	Password  string
	DiscordGo *discordgo.Session
	Guilds    map[string]*discordgo.Guild
}
