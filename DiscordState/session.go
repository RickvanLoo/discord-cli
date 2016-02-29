//Package DiscordState is an abstraction layer that gives proper structs and functions to get and set the current state of the cli server
package DiscordState

import "github.com/Rivalo/discordgo_cli"

//!----- Session -----!//

//NewSession Creates a new Session
func NewSession(Username, Password string) *Session {
	Session := new(Session)
	Session.Username = Username
	Session.Password = Password

	return Session
}

//Start attaches a discordgo listener to the Sessions and fills it.
func (Session *Session) Start() error {
	// TODO: Fill State

	dg, err := discordgo.New(Session.Username, Session.Password)
	if err != nil {
		return err
	}

	// Open the websocket and begin listening.
	dg.Open()

	for _, Guild := range dg.State.Guilds {
		Session.Guilds[Guild.ID] = Guild
	}

	Session.DiscordGo = dg

	return nil
}

//NewState attaches a new state to the Guild inside a Session, and fills it
func (Session *Session) NewState(GuildID string) *State {
	State := new(State)

	//Set Session
	State.Session = Session

	//Set Guild
	State.Guild = Session.Guilds[GuildID]

	//Retrieve Channels
	for _, Channel := range State.Guild.Channels {
		State.Channels[Channel.ID] = Channel
	}

	//Retrieve Members
	for _, Member := range State.Guild.Members {
		State.Members[Member.User.ID] = Member
	}

	return State
}

//Update updates the current Guilds inside the session
func (Session *Session) Update() {
	NewGuildList := make(map[string]*discordgo.Guild)

	for _, Guild := range Session.DiscordGo.State.Guilds {
		NewGuildList[Guild.ID] = Guild
	}

	Session.Guilds = NewGuildList
}
