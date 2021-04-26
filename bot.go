package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token string = os.Getenv("TOKEN")
var user *discordgo.User

func main() {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating Discord client", err)
		return
	}

	if err := discord.Open(); err != nil {
		fmt.Println("Error opening Discord connection")
		return
	}

	defer discord.Close()

	discord.AddHandler(messageHandler)

	user, _ = discord.User("@me")

	fmt.Printf("%s is now connected !", user.Username)

	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == user.ID {
		return
	}

	var messageReceived string = m.Content
	var messageToSend string

	if strings.HasPrefix(messageReceived, "! ping") {
		messageToSend = "pong"
	}

	if strings.HasPrefix(messageReceived, "! quote") {
		var err error
		messageToSend, err = GetRandomQuote()

		if err != nil {
			messageToSend = "Oops, could not load a quote at the moment"
		}
	}

	if messageToSend == "" {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, messageToSend)

	if err != nil {
		fmt.Println("Error sending message")
	}
}

// features
// meme : reddit
// joke : https://sv443.net/jokeapi/v2/
// quote : zenquotes.io
// chatbot : can use facebook messenger api
