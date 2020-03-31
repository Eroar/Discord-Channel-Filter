package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

const token = "NjkzNTQyNjA4NzczMDU0NDg1.XoINdg.DRbXx9iE2xLi8sHrVzeMxPY_y7c"

func main() {
	args := os.Args[1:]
	channelID, messageID := args[0], args[1]

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error:\n", err)
		return
	}

	msg, err := discord.ChannelMessage(channelID, messageID)
	if err != nil {
		fmt.Println("Error:\n", err)
		return
	}
	fmt.Println("Content:\n", msg.Content)

}
