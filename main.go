// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
	"log"

	"github.com/Rivalo/readline"
	"github.com/bwmarrin/discordgo"
)

// Session contains the current settings of the client
type Session struct {
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Guild      *discordgo.Guild   `json:"guild"`
	Channel    *discordgo.Channel `json:"channel"`
	InsertMode bool               `json:"-"`
}

func main() {
	//Initialize Config
	GetConfig()
	CheckState()
	State.InsertMode = false
	Clear()
	Header("V0.1")

	// Connect to Discord
	dg, err := discordgo.New(State.Username, State.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register messageCreate as a callback for the OnMessageCreate event.
	dg.OnMessageCreate = messageCreate

	// Open the websocket and begin listening.
	dg.Open()

	//Print Welcome as a sign that the user has logged in.
	Welcome(dg)

	//SetChannelState
	SetGuildState(dg)
	SetChannelState(dg)

	//Setup stdout logging
	rl, err := readline.NewEx(&readline.Config{
		Prompt:         "> ",
		UniqueEditLine: true,
	})
	if err != nil {
		panic(err)
	}

	defer rl.Close()
	log.SetOutput(rl.Stderr()) // let "log" write to l.Stderr instead of os.Stderr

	//Start Listening
	for {
		line, _ := rl.Readline()

		//QUIT
		if line == ":q" {
			break
		}

		line = ParseForCommands(line, dg)

		if line != "" {
			dg.ChannelMessageSend(State.Channel.ID, line)
		}
	}

	return
}
