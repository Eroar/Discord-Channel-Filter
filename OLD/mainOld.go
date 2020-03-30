package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const token = "NjkzNTQyNjA4NzczMDU0NDg1.Xn-0NA.VRy_X8UIrFYNyTKWJzy6Pzg9ii8"

var musicChannels = [...]string{"music_commands"}
var musicBots = [...]string{"BARD SERWEROWY"}

// func randomReprimand() string {

// }

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	//Create session with token
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error during creating session", err)
		return
	} else {
		fmt.Println("Session created!")
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		fmt.Println("Message", m.Content)
		// channels, err := s.Gu nnels))
		// if err == nil {
		// 	fmt.Println(len(channels))
		// 	for i, channel := range channels {
		// 		fmt.Printf("\nChannel %d:\n\tName: %s\n\tID: %s", i, channel.Name, channel.ID)

		// 	}
		// }
		var channel, _ = s.Channel(m.ChannelID)
		fmt.Println("Channel Name: ", channel.Name)
		if !stringInSlice(channel.Name, musicChannels[:]) {
			if strings.Compare(string([]rune(m.Content)[0]), "!") == 0 {
				s.ChannelMessageSendTTS(m.ChannelID, "Oj, ty hultaju!")
				messagesAfter, err := s.ChannelMessages(m.ChannelID, 0, "", m.ID, "")
				var messages2delete []discordgo.Message

				if err == nil {
					//decides which messages are from bot
					for _, message := range messagesAfter {
						fmt.Println("Message", m.Content)
						for _, role := range musicBots {
							if stringInSlice(role, message.Member.Roles) {
								messages2delete = append(messages2delete, *message)
								break
							}
						}
					}
					fmt.Println("Messages: ", messages2delete)

					for _, message := range messages2delete {
						s.ChannelMessageDelete(m.ChannelID, message.ID)
					}

					s.ChannelMessageDelete(m.ChannelID, m.ID)
				}

			}
		}
	})

	err = discord.Open()
	if err != nil {
		fmt.Println("Error during creating connection")
		return
	}
	//Set status to make the bot online
	discord.UserUpdateStatus(discordgo.StatusOnline)
	fmt.Println("Bot Online!")

	//Keeps the bot running
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
