package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Rivalo/discordgo_cli"
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
	Msg(InfoMsg, "Welcome, %s!\n\n", State.Session.User.Username)
	Msg(InfoMsg, "Guild: %s, Channel: %s\n", State.Guild.Name, State.Channel.Name)
}

//ReceivingMessageParser parses receiving message for mentions, images and MultiLine and returns string array
func ReceivingMessageParser(m *discordgo.Message) []string {
	Message := m.ContentWithMentionsReplaced()

	//Parse images
	for _, Attachment := range m.Attachments {
		Message = Message + " " + Attachment.URL
	}

	// MultiLine comment parsing
	Messages := strings.Split(Message, "\n")

	return Messages
}

//PrintMessages prints amount of Messages to CLI
func PrintMessages(Amount int) {
	for Key, m := range State.Messages {
		if Key >= len(State.Messages)-Amount {
			Messages := ReceivingMessageParser(m)

			for _, Msg := range Messages {
				//log.Printf("> %s > %s\n", UserName(m.Author.Username), Msg)
				MessagePrint(m.Timestamp, m.Author.Username, Msg)

			}
		}
	}
}

//Notify uses Notify-Send from libnotify to send a notification when a mention arrives.
func Notify(m *discordgo.Message) {
	Channel, err := State.Session.DiscordGo.Channel(m.ChannelID)
	if err != nil {
		Msg(ErrorMsg, "(NOT) Channel Error: %s\n", err)
	}
	Guild, err := State.Session.DiscordGo.Guild(Channel.GuildID)
	if err != nil {
		Msg(ErrorMsg, "(NOT) Guild Error: %s\n", err)
	}
	Title := "@" + m.Author.Username + " : " + Guild.Name + "/" + Channel.Name
	cmd := exec.Command("notify-send", Title, m.ContentWithMentionsReplaced())
	err = cmd.Start()
	if err != nil {
		Msg(ErrorMsg, "(NOT) Check if libnotify is installed, or disable notifications.\n")
	}

}

//MessagePrint prints one correctly formatted Message to stdout
func MessagePrint(Time, Username, Content string) {
	UserName := color.New(color.FgGreen).SprintFunc()
	TimeStamp, _ := time.Parse(time.RFC3339, Time)
	LocalTime := TimeStamp.Local().Format("2006/01/02 15:04:05")
	log.SetFlags(0)
	log.Printf("%s > %s > %s\n", LocalTime, UserName(Username), Content)
	log.SetFlags(log.LstdFlags)
}
