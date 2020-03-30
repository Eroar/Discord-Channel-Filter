package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const token = "NjkzNTQyNjA4NzczMDU0NDg1.Xn-0NA.VRy_X8UIrFYNyTKWJzy6Pzg9ii8"

var musicChannels = [...]string{"music_commands"}
var musicBotsIDs[]string = {""}

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
		if err != nil {
			fmt.Printf("\nError during handling a message %s", m.ID)
		}
		if !stringInSlice(m.ChannelID, musicChannels) {

		}
		fmt.Println("Channel Name: ", channel.Name)
		var id = m.Member.User.ID
		fmt.Printf("\nMessage on %s\nfrom: %s\ncontent: %s", channel.Name, id, m.Content)
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
