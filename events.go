package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

// This function will be called (due to above assignment) every time a new
// message is created on any channel that the autenticated user has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	UserName := color.New(color.FgGreen).SprintFunc()

	if State.InsertMode {
		if m.ChannelID == State.Channel.ID {
			log.Printf("> %s > %s\n", UserName(m.Author.Username), m.Content)
		}
	}

}
