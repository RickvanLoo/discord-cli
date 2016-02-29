package DiscordState

//SetChannel sets the channel of the current State
func (State *State) SetChannel(ID string) {
	State.Channel = State.Channels[ID]
}

//UpdateChannels updates the channels inside the current State
func (State *State) UpdateChannels() {

}

//AddMember adds Member to State
func (State *State) AddMember(ID string) error {
	Member, err := State.Session.DiscordGo.GuildMember(State.Guild.ID, ID)
	if err != nil {
		return err
	}

	State.Members[ID] = Member
	return nil
}

//DelMember deletes Member from State
func (State *State) DelMember(ID string) {
	delete(State.Members, ID)
}

//RetrieveMessages retrieves last N messages inside channel
func (State *State) RetrieveMessages(amount int) error {
	Messages, err := State.Session.DiscordGo.ChannelMessages(State.Channel.ID, amount, "", "")
	if err != nil {
		return err
	}

	for _, Message := range Messages {
		State.Messages[Message.ID] = Message
	}

	return nil
}

//AddMessage adds a message to the State
func (State *State) AddMessage(ID string) {

}

//EditMessage edits a message inside the State
func (State *State) EditMessage(ID string) {

}

//DelMessage deletes a message inside the State
func (State *State) DelMessage(ID string) {

}
