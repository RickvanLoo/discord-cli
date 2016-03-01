package main

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string) string {
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

	return line
}

//SelectGuild selects a new Guild
func SelectGuild() {
	State.Enabled = false
	SelectGuildMenu()
	SelectChannelMenu()
	State.Enabled = true
	Clear()
	Header()
}

//SelectChannel selects a new Channel
func SelectChannel() {
	State.Enabled = false
	SelectChannelMenu()
	State.Enabled = true
	Clear()
	Header()
}
