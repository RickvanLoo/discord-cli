// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/Rivalo/discord-cli/DiscordState"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
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
	Session := DiscordState.NewSession(Config.Username, Config.Password) //Please don't abuse
	err := Session.Start()
	if err != nil {
		log.Println("Session Failed")
		log.Fatalln(err)
	}

	//NewState
	State = Session.NewState("148824204219777024")
	State.SetChannel("148824204219777024")

	State.Session.DiscordGo.AddHandler(newMessage)

	//Print Welcome as a sign that the user has logged in.
	user, _ := State.Session.DiscordGo.User("@me")
	Msg(InfoMsg, "Welcome, %s!\n\n", user.Username)

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

		if line != "" {
			State.Session.DiscordGo.ChannelMessageSend(State.Channel.ID, line)
		}
	}

	return
}

//Msg is a composition of Color.New printf functions
func Msg(MsgType, format string, a ...interface{}) {

	// TODO: Add support for changing color by configuration

	Error := color.New(color.FgRed, color.Bold)
	Info := color.New(color.FgYellow, color.Bold)
	Head := color.New(color.FgCyan, color.Bold)
	Text := color.New(color.FgWhite)

	switch MsgType {
	case "Error":
		Error.Printf(format, a...)
	case "Info":
		Info.Printf(format, a...)
	case "Head":
		Head.Printf(format, a...)
	case "Text":
		Text.Printf(format, a...)
	default:
		Text.Printf(format, a...)
	}
}

//Clear clears the terminal => This barely works, please fix
func Clear() {

	// TODO: ADD support for multiple operating systems and terminals. Linux = clear, Windows = cls, have to do research for OSX and BSD.

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
