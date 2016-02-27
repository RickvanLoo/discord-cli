package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Rivalo/discordgo_cli"
)

//THIS FILE IS A COMPLETE MESS, IT BARELY WORKS
//PLEASE FIX

//SetChannelState sets the Channel inside the State
func SetChannelState(dg *discordgo.Session) {
	State.InsertMode = false

	guild := State.Guild
	Msg(InfoMsg, "Select a Channel:\n")
	for key, channel := range guild.Channels {
		if channel.Type == "text" {
			Msg(TextMsg, "%d:%s\n", key, channel.Name)
		}
	}
	Msg(InfoMsg, "b:Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)
	GuildID, _ := strconv.Atoi(response)

	if response == "b" {
		SetGuildState(dg)
		return
	}

	for guild.Channels[GuildID].Type != "text" {
		Msg(ErrorMsg, "That's a voice channel, you know this is a CLI right?\n")
		Msg(InfoMsg, "Select a Channel:\n")
		fmt.Scanf("%s\n", &response)
	}

	GuildID, _ = strconv.Atoi(response)
	State.Channel = guild.Channels[GuildID]

	Clear()

	State.InsertMode = true
}

//SetGuildState sets the Guild inside the State
func SetGuildState(dg *discordgo.Session) {
	State.InsertMode = false
	Guilds, _ := dg.UserGuilds()
	Msg(InfoMsg, "Select a Guild:\n")

	for key, guild := range Guilds {
		Msg(TextMsg, "%d:%s\n", key, guild.Name)

	}
	printExtraOptions()

	var response string
	fmt.Scanf("%s\n", &response)

	switch response {
	case "n":
		ServJoin(dg)
	case "d":
		ServLeave(dg)
	case "o":
		ServJoinDiscordCli(dg)
	case "q":
		os.Exit(0)
	default:
		GuildID, _ := strconv.Atoi(response)
		State.Guild, _ = dg.Guild(Guilds[GuildID].ID)
		Clear()
		SetChannelState(dg)

	}
	return
}

func printExtraOptions() {
	Msg(TextMsg, "\nn:Join Server\n")
	Msg(TextMsg, "d:Leave Server.\n")
	Msg(TextMsg, "o:Join discord-cli Server\n")
	Msg(TextMsg, "q:Quit discord-cli\n")
}

//ServJoinDiscordCli joins the official discord-cli server
func ServJoinDiscordCli(dg *discordgo.Session) {
	Clear()

	Msg(InfoMsg, "Join official discord-cli server? [Y/n]\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if strings.ToUpper(response) == "" || strings.ToUpper(response) == "Y" {
		dg.InviteAccept("0pXWCo5RQbVuFHDM")
		SetGuildState(dg)
	} else {
		SetGuildState(dg)
	}

	return
}

//ServJoin joins a Server based on invite
func ServJoin(dg *discordgo.Session) {
	Clear()
	Msg(InfoMsg, "Enter invite code to enter server: (empty to go back)\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "" {
		SetGuildState(dg)
		return
	}

	Invite, err := dg.Invite(response)
	if err != nil {
		fmt.Println(err)
		SetGuildState(dg)
		return
	}

	Msg(InfoMsg, "Join %s [Y/n]?", Invite.Guild.Name)

	var confirm string
	fmt.Scanf("%s\n", &confirm)
	if strings.ToUpper(confirm) == "" || strings.ToUpper(confirm) == "Y" {
		dg.InviteAccept(response)
		SetGuildState(dg)
	} else {
		SetGuildState(dg)
	}
}

//ServLeave lists servers, and leaves them based on input
func ServLeave(dg *discordgo.Session) {
	Clear()
	Guilds, _ := dg.UserGuilds()
	Msg(InfoMsg, "Leave a server?\n")

	for key, guild := range Guilds {
		Msg(TextMsg, "%d:%s\n", key, guild.Name)
	}
	Msg(InfoMsg, "b:Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		SetGuildState(dg)
		return
	}
	GuildID, _ := strconv.Atoi(response)

	Msg(InfoMsg, "Leave %s [Y/n]?", Guilds[GuildID].Name)

	var confirm string
	fmt.Scanf("%s\n", &confirm)
	if strings.ToUpper(confirm) == "" || strings.ToUpper(confirm) == "Y" {
		dg.GuildDelete(Guilds[GuildID].ID)
		SetGuildState(dg)
	} else {
		SetGuildState(dg)
	}

}
