package zlyden

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "Say Hi" reply with "Hello World!"
	if strings.EqualFold(m.Content, "Say Hi") {
		s.ChannelMessageSend(m.ChannelID, "Hello World!")
	}

	// If the message is "Say Goodbye" reply with "Goodbye World!"
	if strings.EqualFold(m.Content, "Say Goodbye") {
		s.ChannelMessageSend(m.ChannelID, "Goodbye World!")
		s.Close()
	}
}
