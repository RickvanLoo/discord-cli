package main

import (
	"strconv"
	"strings"
)

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string) string {
	//One Key Commands
	switch line {
	case ":g":
		SelectGuild()
		line = ""
	case ":c":
		SelectChannel()
		line = ""
	default:
		// Nothing
	}

	//Argument Commands
	if strings.HasPrefix(line, ":m") {
		AmountStr := strings.Split(line, " ")
		if len(AmountStr) < 2 {
			Msg(ErrorMsg, "[:m] No Arguments \n")
			return ""
		}

		Amount, err := strconv.Atoi(AmountStr[1])
		if err != nil {
			Msg(ErrorMsg, "[:m] Argument Error: %s \n", err)
			return ""
		}

		Msg(InfoMsg, "Printing last %d messages!\n", Amount)
		State.RetrieveMessages(Amount)
		PrintMessages(Amount)
		line = ""
	}

	return line
}

//SelectGuild selects a new Guild
func SelectGuild() {
	State.Enabled = false
	SelectGuildMenu()
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}

//SelectChannel selects a new Channel
func SelectChannel() {
	State.Enabled = false
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}
