package main

import "github.com/bwmarrin/discordgo"

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string, dg *discordgo.Session) string {
	switch line {
	case ":d":
		SetGuildState(dg)
		line = ""
	case ":c":
		SetChannelState(dg)
		line = ""
	default:
		// Nothing

	}

	return line
}
