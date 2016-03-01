package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
)

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

//Header simply prints a header containing state/session information
func Header() {
	user, _ := State.Session.DiscordGo.User("@me")
	Msg(InfoMsg, "Welcome, %s!\n\n", user.Username)
	Msg(InfoMsg, "Guild: %s, Channel: %s\n", State.Guild.Name, State.Channel.Name)
}
