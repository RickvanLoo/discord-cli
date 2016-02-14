package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

//THIS FILE IS A COMPLETE MESS, IT BARELY WORKS
//PLEASE FIX

func Header(version string) {
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("discord-cli - version: %s\n\n", version)
}

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Welcome(dg *discordgo.Session) {
	d := color.New(color.FgYellow, color.Bold)
	d.Printf("Listening!\n\n")

	user, _ := dg.User("@me")
	d.Printf("Welcome, %s!\n\n", user.Username)
}

func SetChannelState(dg *discordgo.Session) {
	State.InsertMode = false

	Guilds, _ := dg.UserGuilds()
	d := color.New(color.FgYellow, color.Bold)
	d.Printf("Select a Guild:\n")

	for key, guild := range Guilds {
		fmt.Printf("%d:%s\n", key, guild.Name)
	}

	var response int
	fmt.Scanf("%d\n", &response)

	guild, _ := dg.Guild(Guilds[response].ID)

	d.Printf("Select a Channel:\n")
	for key, channel := range guild.Channels {
		if channel.Type == "text" {
			fmt.Printf("%d:%s\n", key, channel.Name)
		}
	}

	fmt.Scanf("%d\n", &response)
	for guild.Channels[response].Type != "text" {
		Error := color.New(color.FgRed, color.Bold)
		Error.Printf("That's a voice channel, you know this is a CLI right?\n")
		d.Printf("Select a Channel:\n")
		fmt.Scanf("%d\n", &response)
	}

	State.Channel = guild.Channels[response]
	State.Guild = guild

	Clear()

	State.InsertMode = true
}
