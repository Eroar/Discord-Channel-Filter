package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

var musicChannelsIds []string
var musicBotsIds []string

//Config containis for parsing json
type Config struct {
	Token            string
	MusicChannelsIds []string
	MusicBotsIds     []string
}

func getSettings() Config {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print(err)
	}
	if !json.Valid(configFile) {
		fmt.Println("Json file not valid")
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isMusicCommand(m *discordgo.Message) bool {
	if len(m.Content) > 0 {
		if strings.Compare(string([]rune(m.Content)[0]), "!") == 0 {
			return true
		}
	}
	if stringInSlice(m.Author.ID, musicBotsIds) {
		return true
	}
	return false
}

func main() {
	//Create session with token
	settings := getSettings()
	fmt.Println(settings)
	token = settings.Token
	musicBotsIds = settings.MusicBotsIds
	musicChannelsIds := settings.MusicChannelsIds

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error during creating session", err)
		return
	} else {
		fmt.Println("Session created!")
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		var channel, err = s.Channel(m.ChannelID)
		if err != nil {
			fmt.Printf("\nError during handling a message %s:\n", m.ID)
			fmt.Print(err)
			fmt.Println("")
		}
		fmt.Printf("\nMessage on %s\nfrom: %s\ncontent: %s\n", channel.Name, m.Author.ID, m.Content)
		discordgo.Message
		// If channel is not a music channel
		if !stringInSlice(m.ChannelID, musicChannelsIds) {
			if isMusicCommand(m.Message) {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
			}
		}
	})

	err = discord.Open()
	if err != nil {
		fmt.Println("Error during creating connection")
		fmt.Print(err)
		fmt.Println("")
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
