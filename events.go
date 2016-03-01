package main

import (
	"log"

	"github.com/Rivalo/discordgo_cli"
	"github.com/fatih/color"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Print message to stdout.
	UserName := color.New(color.FgGreen).SprintFunc()
	if m.ChannelID == State.Channel.ID {
		log.Printf("> %s > %s\n", UserName(m.Author.Username), m.Content)
	}
}
