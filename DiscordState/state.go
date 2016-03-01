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

//RetrieveMessages retrieves last N messages inside channel
func (State *State) RetrieveMessages(Amount int) error {
	Messages, err := State.Session.DiscordGo.ChannelMessages(State.Channel.ID, Amount, "", "")
	if err != nil {
		return err
	}

	NewMessageList := make(map[string]*discordgo.Message)

	for _, Message := range Messages {
		NewMessageList[Message.ID] = Message
	}

	State.Messages = NewMessageList

	return nil
}

//AddMessage adds a message to the State
func (State *State) AddMessage(Message *discordgo.Message) {
	State.Messages[Message.ID] = Message
}

//EditMessage edits a message inside the State
func (State *State) EditMessage(Message *discordgo.Message) {
	State.Messages[Message.ID] = Message
}

//DelMessage deletes a message inside the State
func (State *State) DelMessage(Message *discordgo.Message) {
	delete(State.Messages, Message.ID)
}
