package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

//MessageCreate is a handler for discordgo.MessageCreate
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Message event")
	channels, err := s.GuildChannels(m.GuildID)
	fmt.Println(len(channels))
	if err == nil {
		fmt.Println(len(channels))
		for i, channel := range channels {
			fmt.Printf("\nChannel %d:\n\tName: %s\n\tID: %s", i, channel.Name, channel.ID)

		}
	}
	if m.Content == "!HELLO" {
		s.ChannelMessageSend(m.ChannelID, "Hello, I'm a bot")
	}
}
