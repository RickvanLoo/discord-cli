package DiscordState

import "github.com/Rivalo/discordgo_cli"

//SetChannel sets the channel of the current State
func (State *State) SetChannel(ID string) {
	for _, Channel := range State.Channels {
		if Channel.ID == ID {
			State.Channel = Channel
		}
	}
}

//AddMember adds Member to State
func (State *State) AddMember(Member *discordgo.Member) {
	State.Members[Member.User.ID] = Member
}

//DelMember deletes Member from State
func (State *State) DelMember(Member *discordgo.Member) {
	delete(State.Members, Member.User.ID)
}

//AddMessage adds Message to State
func (State *State) AddMessage(Message *discordgo.Message) {
	//Do not add if Amount <= 0
	if State.MessageAmount <= 0 {
		return
	}

	//Remove First Message if next message is going to increase length past MessageAmount
	if len(State.Messages) == State.MessageAmount {
		State.Messages = append(State.Messages[:0], State.Messages[1:]...)
	}

	State.Messages = append(State.Messages, Message)
}

//EditMessage edits Message inside State
func (State *State) EditMessage(Message *discordgo.Message) {
	for Index, StateMessage := range State.Messages {
		if StateMessage.ID == Message.ID {
			State.Messages[Index] = Message
		}
	}
}

//DelMessage deletes Message from State
func (State *State) DelMessage(Message *discordgo.Message) {
	for Index, StateMessage := range State.Messages {
		if StateMessage.ID == Message.ID {
			State.Messages = append(State.Messages[:Index], State.Messages[Index+1:]...)
		}
	}
}

//RetrieveMessages retrieves last N Messages and puts it in state
func (State *State) RetrieveMessages(Amount int) error {
	Messages, err := State.Session.DiscordGo.ChannelMessages(State.Channel.ID, Amount, "", "")
	if err != nil {
		return err
	}

	//Reverse insert Messages
	for i := 0; i < len(Messages); i++ {
		State.AddMessage(Messages[len(Messages)-i-1])
	}

	return nil
}
