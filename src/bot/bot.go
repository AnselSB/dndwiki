package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"wikiBot/src/rest"

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

// TODO: for spell check that an argument is given
// this will be the main event handler function, most of what this bot will do is recieve and send messages, it doesn't need to do anything super fancy
// once I get this out of testing phase I'll probably make a new package that will hold the different logic for each functionality, one for magic searching
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// first check to see we're responding to a user and not ourself
	if message.Author.ID == discord.State.User.ID {
		return
	}
	// split the message and compare to see if the first substring in the message matches the command in the switch statement
	splitMsg := strings.Split(message.Content, " ")

	switch splitMsg[0] {
	case "!hello":
		discord.ChannelMessageSend(message.ChannelID, "My balls itch")
	case "!spell":
		if len(splitMsg) < 2 {
			discord.ChannelMessageSend(message.ChannelID, "Please provide a spell you wish to search for")
			return
		}
		embed, err := rest.GetSpell(splitMsg[1])
		if err != nil {
			discord.ChannelMessageSend(message.ChannelID, "Error fetching spell, make sure to check spelling")
			return
		}
		discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	}
}
