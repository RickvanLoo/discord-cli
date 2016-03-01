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

	//Retrieve GuildID's from current User
	UserGuilds, err := dg.UserGuilds()
	if err != nil {
		return err
	}

	Session.Guilds = UserGuilds

	Session.DiscordGo = dg

	return nil
}

//NewState attaches a new state to the Guild inside a Session, and fills it.
func (Session *Session) NewState(GuildID string) (*State, error) {
	State := new(State)

	//Disable Event Handling
	State.Enabled = false

	//Set Session
	State.Session = Session

	//Set Guild
	for _, guildID := range Session.Guilds {
		if guildID.ID == GuildID {
			Guild, err := State.Session.DiscordGo.Guild(guildID.ID)
			if err != nil {
				return nil, err
			}

			State.Guild = Guild
		}
	}

	//Retrieve Channels

	State.Channels = State.Guild.Channels

	//Retrieve Members

	State.Members = make(map[string]*discordgo.Member)

	for _, Member := range State.Guild.Members {
		State.Members[Member.User.ID] = Member
	}

	return State, nil
}

//Update updates the session, this reloads the Guild list
func (Session *Session) Update() error {
	UserGuilds, err := Session.DiscordGo.UserGuilds()
	if err != nil {
		return err
	}

	Session.Guilds = UserGuilds
	return nil
}
