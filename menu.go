package main

import (
	"fmt"
	"log"
)

//SelectGuild is a menu item that creates a new State on basis of Guild selection
func SelectGuild() {
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

//SelectChannel is a menu item that sets the current channel
func SelectChannel() {
	Msg(InfoMsg, "Select a Channel:\n")

	SelectMap := make(map[int]string)

	for id, channel := range State.Channels {
		SelectMap[id] = channel.ID
		Msg(TextMsg, "[%d] %s\n", id, channel.Name)
	}

	var response int
	fmt.Scanf("%d\n", &response)

	State.SetChannel(SelectMap[response])
}
