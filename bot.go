package main

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token string = os.Getenv("TOKEN")
var user *discordgo.User

func main() {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalf("Error creating Discord client: %s", err)
	}

	if err := discord.Open(); err != nil {
		log.Fatalf("Error opening Discord connection: %s", err)
	}

	defer discord.Close()

	discord.AddHandler(messageHandler)

	user, _ = discord.User("@me")

	log.Printf("%s is now connected", user.Username)

	// this prevents the function from terminating and keeps the server listening
	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == user.ID {
		return
	}

	var messageReceived string = m.Content
	var messageToSend string
	var err error
	// figure out the type of command
	// generate the message
	// send the message

	// ping handler
	if strings.HasPrefix(messageReceived, "! ping") {
		messageToSend = "pong"
	}

	// quote handler
	if strings.HasPrefix(messageReceived, "! quote") {
		messageToSend, err = GetRandomQuote()

		if err != nil {
			messageToSend = "Oops, could not load a quote at the moment"
		}
	}

	// joke handler
	if strings.HasPrefix(messageReceived, "! joke") {
		var category string

		words := strings.Fields(messageReceived)

		if len(words) == 2 {
			category = ""
		} else {
			category = words[2]
		}

		messageToSend, err = GetRandomJoke(category)

		if err != nil {
			messageToSend = "Oops, could not load a joke at the moment. Check if your command is valid"
		}
	}

	if messageToSend == "" {
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, messageToSend)

	if err != nil {
		log.Printf("Error sending message '%s'", messageToSend)
	}
}
