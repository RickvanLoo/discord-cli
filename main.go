// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"log"

	"github.com/Rivalo/discord-cli/DiscordState"
	"github.com/chzyer/readline"
)

//Global Message Types
const (
	ErrorMsg  = "Error"
	InfoMsg   = "Info"
	HeaderMsg = "Head"
	TextMsg   = "Text"
)

//Version is current version const
const Version = "v0.3.0-DEVELOP"

//Session is global Session
var Session *DiscordState.Session

//State is global State
var State *DiscordState.State

//MsgType is a string containing global message type
type MsgType string

func main() {
	//Initialize Config
	GetConfig()
	CheckState()
	Clear()
	Msg(HeaderMsg, "discord-cli - version: %s\n\n", Version)

	//NewSession
	Session = DiscordState.NewSession(Config.Username, Config.Password) //Please don't abuse
	err := Session.Start()
	if err != nil {
		log.Println("Session Failed")
		log.Fatalln(err)
	}

	//Attach New Window
	InitWindow()

	//Attach Even Handlers
	State.Session.DiscordGo.AddHandler(newMessage)

	//Print Welcome as a sign that the user has logged in.
	user, _ := State.Session.DiscordGo.User("@me")
	Msg(InfoMsg, "Welcome, %s!\n\n", user.Username)
	Msg(InfoMsg, "Guild: %s, Channel: %s\n", State.Guild.Name, State.Channel.Name)

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

		line = ParseForCommands(line)

		if line != "" {
			State.Session.DiscordGo.ChannelMessageSend(State.Channel.ID, line)
		}
	}

	return
}

//InitWindow creates a New CLI Window
func InitWindow() {
	SelectGuildMenu()
	SelectChannelMenu()
	State.Enabled = true
	Clear()
	State.RetrieveMessages(10)
}
