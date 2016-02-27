// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Rivalo/discordgo_cli"
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
const Version = "v0.2.0"

//MsgType is a string containing global message type
type MsgType string

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
	Msg(HeaderMsg, "discord-cli - version: %s\n\n", Version)

	// Connect to Discord
	dg, err := discordgo.New(State.Username, State.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register messageCreate as a callback for the OnMessageCreate event.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	dg.Open()

	//Print Welcome as a sign that the user has logged in.
	user, _ := dg.User("@me")
	Msg(InfoMsg, "Welcome, %s!\n\n", user.Username)

	//SetChannelState
	SetGuildState(dg)

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

//Msg is a composition of Color.New printf functions
func Msg(MsgType, format string, a ...interface{}) {

	// TODO: Add support for changing color by configuration

	Error := color.New(color.FgRed, color.Bold)
	Info := color.New(color.FgYellow, color.Bold)
	Head := color.New(color.FgCyan, color.Bold)
	Text := color.New(color.FgWhite)

	//Testing Colors
	//Info := color.New(color.FgMagenta, color.CrossedOut)
	//Head := color.New(color.FgMagenta, color.CrossedOut)
	//Text := color.New(color.FgMagenta, color.CrossedOut)

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
