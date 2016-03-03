package main

import (
	"log"
	"strings"

	"github.com/Rivalo/discordgo_cli"
	"github.com/fatih/color"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Do nothing when State is disabled
	if !State.Enabled {
		return
	}

	// Print message to stdout.
	UserName := color.New(color.FgGreen).SprintFunc()
	if m.ChannelID == State.Channel.ID {
		//State.AddMessage(m.Message) //BROKEN
		Message := m.ContentWithMentionsReplaced()

		//Parse images
		for _, Attachment := range m.Attachments {
			Message = Message + " " + Attachment.URL
		}

		// MultiLine comment parsing
		Messages := strings.Split(Message, "\n")

		for _, Msg := range Messages {
			log.Printf("> %s > %s\n", UserName(m.Author.Username), Msg)
		}
	}
}
