//Package state gives proper structs and functions to get and set the current state of the cli server
package state

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

//New creates a New state
func New(Username, Password string) *State {
	NewState := new(State)
	NewState.Username = Username
	NewState.Password = Password

	return NewState
}

//Start attaches a discordgo listener to the first state.
func Start(State *State) {
	dg, err := discordgo.New(State.Username, State.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open the websocket and begin listening.
	dg.Open()
}
