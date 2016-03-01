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

	dg, err := discordgo.New(Session.Username, Session.Password)
	if err != nil {
		return err
	}

	// Open the websocket and begin listening.
	dg.Open()

	Session.Guilds = make(map[string]*discordgo.Guild)

	//Retrieve GuildID's from current User
	UserGuilds, err := dg.UserGuilds()
	if err != nil {
		return err
	}

	//Retrieve Guilds from GuildIDs
	for _, UserGuild := range UserGuilds {
		Guild, err := dg.Guild(UserGuild.ID)
		if err != nil {
			return err
		}
		Session.Guilds[Guild.ID] = Guild
	}

	Session.DiscordGo = dg

	return nil
}

//NewState attaches a new state to the Guild inside a Session, and fills it.
func (Session *Session) NewState(GuildID string) *State {
	State := new(State)

	//Disable Event Handling
	State.Enabled = false

	//Set Session
	State.Session = Session

	//Set Guild
	State.Guild = Session.Guilds[GuildID]

	//Retrieve Channels

	State.Channels = make(map[string]*discordgo.Channel)

	for _, Channel := range State.Guild.Channels {
		State.Channels[Channel.ID] = Channel
	}

	//Retrieve Members

	State.Members = make(map[string]*discordgo.Member)

	for _, Member := range State.Guild.Members {
		State.Members[Member.User.ID] = Member
	}

	return State
}

//Update does a full update for the Guilds inside the State
func (Session *Session) Update() {
	NewGuildList := make(map[string]*discordgo.Guild)

	for _, Guild := range Session.DiscordGo.State.Guilds {
		NewGuildList[Guild.ID] = Guild
	}

	Session.Guilds = NewGuildList
}
