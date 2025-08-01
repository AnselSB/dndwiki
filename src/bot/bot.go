package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	// first make the session with discord
	fmt.Printf("BotTOKEN: %v\n", BotToken)
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add an event handler for the bot, this event handler is a function that you pass in
	discord.AddHandler(newMessage) // the function signature will determine what event handler you add

	// open the session and defer a close to make sure things are closed properly
	discord.Open()
	defer discord.Close()

	// keep the bot running until there is an interruption by the user (which will most likely be me, this'll mean the bot will only be running when I decide it to be running, not really interested in hosting this on the cloud, too expensive for what this is)
	fmt.Println("wiki bot is now running, press ctrl c to stop execution")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

// this will be the main event handler function, most of what this bot will do is recieve and send messages, it doesn't need to do anything super fancy
// once I get this out of testing phase I'll probably make a new package that will hold the different logic for each functionality, one for magic searching
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// first check to see we're responding to a user and not ourself
	if message.Author.ID == discord.State.User.ID {
		return
	}
	// respond to the user message if it contains the words !hello or !magic, don't love contains here I might change this so that it will only trigger if it's at the beginning of the string, could do it through regex but we'll see
	switch {
	case strings.Contains(message.Content, "!hello"):
		discord.ChannelMessageSend(message.ChannelID, "My balls itch")
	case strings.Contains(message.Content, "!spell"):
		discord.ChannelMessageSend(message.ChannelID, "This command doesn't do anything yet because Ansel is a lazy fuck")
	}
}
