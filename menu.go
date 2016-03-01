package main

import (
	"fmt"
	"log"
)

//SelectGuildMenu is a menu item that creates a new State on basis of Guild selection
func SelectGuildMenu() {
	var err error

	Msg(InfoMsg, "Select a Guild:\n")

	SelectMap := make(map[int]string)

	for id, guild := range Session.Guilds {
		SelectMap[id] = guild.ID
		Msg(TextMsg, "[%d] %s\n", id, guild.Name)
	}

	var response int
	fmt.Scanf("%d\n", &response)

	State, err = Session.NewState(SelectMap[response])
	if err != nil {
		log.Fatal(err)
	}
}

//SelectChannelMenu is a menu item that sets the current channel
func SelectChannelMenu() {
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

	var response int
	fmt.Scanf("%d\n", &response)

	State.SetChannel(SelectMap[response])
}
