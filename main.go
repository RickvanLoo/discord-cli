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
	Header("V0.2.0")

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
	Welcome(dg)

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

//Header prints a Cyan header to the TERM containing the program title and its version
func Header(version string) {
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("discord-cli - version: %s\n\n", version)
}

//Clear clears the terminal => This barely works, please fix
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//Welcome sends an acknowledge to the terminal that it is listening, and prints the current Username
func Welcome(dg *discordgo.Session) {
	d := color.New(color.FgYellow, color.Bold)
	d.Printf("Listening!\n\n")

	user, _ := dg.User("@me")
	d.Printf("Welcome, %s!\n\n", user.Username)
}
