package main

import (
	"fmt"
	"log"
	"strconv"
)

//SelectGuildMenu is a menu item that creates a new State on basis of Guild selection
func SelectGuildMenu() {
	var err error

Start:

	Msg(InfoMsg, "Select a Guild:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, guild := range Session.Guilds {
		SelectMap[SelectID] = guild.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, guild.Name)
		SelectID++
	}
	Msg(TextMsg, "[b] Extra Options\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		ExtraGuildMenuOptions()
		goto Start
	}

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(GU) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(GU) Error: ID is out of bounds\n")
		goto Start
	}

	State, err = Session.NewState(SelectMap[ResponseInteger])
	if err != nil {
		log.Fatal(err)
	}
}

//SelectChannelMenu is a menu item that sets the current channel
func SelectChannelMenu() {
Start:
	Msg(InfoMsg, "Select a Channel:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, channel := range State.Channels {
		if channel.Type == "text" {
			SelectMap[SelectID] = channel.ID
			Msg(TextMsg, "[%d] %s\n", SelectID, channel.Name)
			SelectID++
		}
	}
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		SelectGuildMenu()
		goto Start
	}

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(CH) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(CH) Error: ID is out of bound\n")
		goto Start
	}

	State.SetChannel(SelectMap[ResponseInteger])
}

//ExtraGuildMenuOptions prints and handles extra options for SelectGuildMenu
func ExtraGuildMenuOptions() {
Start:
	Msg(InfoMsg, "Extra Options:\n")
	Msg(TextMsg, "[n] Join New Server\n")
	Msg(TextMsg, "[d] Leave Server\n")
	Msg(TextMsg, "[o] Join Official discord-cli Server\n")
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	switch response {
	case "n":
	New:
		Msg(TextMsg, "Please input invite number ([b] back):\n")
		fmt.Scanf("%s\n", &response)
		if response == "b" {
			goto Start
		}
		Invite, err := Session.DiscordGo.Invite(response)
		if err != nil {
			Msg(ErrorMsg, "Invalid Invite\n")
			goto New
		}
		Msg(TextMsg, "Join %s ? [y/n]:\n", Invite.Guild.Name)
		fmt.Scanf("%s\n", &response)
		if response == "y" {
			Session.DiscordGo.InviteAccept(Invite.Code)
			err := Session.Update()
			if err != nil {
				Msg(ErrorMsg, "Session Update Failed: %s\n", err)
			}
		} else {
			goto Start
		}
	case "o":
		_, err := Session.DiscordGo.InviteAccept("0pXWCo5RQbVuFHDM")
		if err != nil {
			Msg(ErrorMsg, "Joining Official discord-cli Server failed\n")
			goto Start
		}
		Msg(InfoMsg, "Joined Official discord-cli Server!\n")
	case "d":
		LeaveServerMenu()
		goto Start
	default:
		return
	}

	return
}

//LeaveServerMenu is a copy of SelectGuildMenu that leaves instead of selects
func LeaveServerMenu() {
	var err error

Start:

	Msg(InfoMsg, "Leave a Guild:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, guild := range Session.Guilds {
		SelectMap[SelectID] = guild.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, guild.Name)
		SelectID++
	}
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		return
	}

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(GUD) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(GUD) Error: ID is out of bounds\n")
		goto Start
	}

	Guild, err := Session.DiscordGo.Guild(SelectMap[ResponseInteger])
	if err != nil {
		Msg(ErrorMsg, "(GUD) Unknown Error: %s\n", err)
		goto Start
	}

	Msg(TextMsg, "Leave %s ? [y/n]:\n", Guild.Name)
	fmt.Scanf("%s\n", &response)
	if response == "y" {
		Session.DiscordGo.GuildLeave(Guild.ID)
		err := Session.Update()
		if err != nil {
			Msg(ErrorMsg, "Session Update Failed: %s\n", err)
		}
	} else {
		goto Start
	}

}
