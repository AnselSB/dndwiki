package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"wikiBot/src/cache"
	"wikiBot/src/rest"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

const (
	spellVal = "1"
	itemVal  = "2"
)

// gonna try and set up the slash command for hello
var (
	dmPermission                   = true
	defaultMemberPermissions int64 = discordgo.PermissionAll
	adminUserPermissions     int64 = discordgo.PermissionAdministrator
	commands                       = []*discordgo.ApplicationCommand{
		{
			Name:                     "hello",
			Description:              "Say hello to wikibot",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
		},
		{
			Name:                     "spell",
			Description:              "search for a spell description in the format spell-name",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "spell-name",
					Description: "Spell Name",
					Required:    true,
				},
			},
		},
		{
			Name:                     "cache",
			Description:              "create a new cache directory, generally only for development purposes",
			DefaultMemberPermissions: &adminUserPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "cache-name",
					Description: "cache Name",
					Required:    true,
				},
			},
		},
		{
			Name:                     "mitem",
			Description:              "search for a magic item description in the format item-name",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "item-name",
					Description: "Item Name",
					Required:    true,
				},
			},
		},
		{
			Name:                     "add",
			Description:              "Add a new item to the database",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "add-option",
					Description: "What type of item would you like to add?",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "spell",
							Value: "1",
						},
						{
							Name:  "magic item",
							Value: "2",
						},
					},
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"hello": hello,
		"spell": spell,
		"cache": makeCache,
		"mitem": mitem,
		"add":   add,
	}
)

func add(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options[0].StringValue()
	fmt.Printf("OPTION NAME: %v\n", options)
	switch options {
	case spellVal:
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: spellModal(i),
		})
		if err != nil {
			fmt.Printf("HUGE ERROR: %v\n", err)
		}
	case itemVal:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "magic item was chosen",
			},
		})
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something horribly wrong has gone down here",
			},
		})
	}
}
func mitem(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) < 1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "An item name was not supplied",
			},
		})
	}
	itemArg := options[0]
	embed, err := rest.GetMagicItem(itemArg.StringValue())
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Unable to find item '%v'", itemArg.StringValue()),
			},
		})
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func makeCache(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	cmdMsg := cache.SetDirectory(options[0].StringValue())
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: cmdMsg,
		},
	})
}
func hello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "My balls itch",
		},
	})
}

func spell(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// first get the argument
	options := i.ApplicationCommandData().Options
	if len(options) < 1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "A spell was not supplied",
			},
		})
	}
	spellArg := options[0]
	embed, err := rest.GetSpell(spellArg.StringValue())
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Unable to find spell '%v'", spellArg.StringValue()),
			},
		})
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func Run() {
	// first make the session with discord
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add handler which will be called when ready
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
	})

	// add handler for slash command handling
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// open the session and defer a close to make sure things are closed properly, also add the necessary commands
	discord.Open()
	fmt.Println("Adding slash commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for index, value := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", value)

		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v", value.Name, err)
		}
		registeredCommands[index] = cmd
		fmt.Printf("Added %v\n", value.Name)
	}
	defer discord.Close()

	// keep the bot running until there is an interruption by the user (which will most likely be me, this'll mean the bot will only be running when I decide it to be running, not really interested in hosting this on the cloud, too expensive for what this is)
	fmt.Println("wiki bot is now running, press ctrl c to stop execution")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Removing commands...")

	for _, v := range registeredCommands {
		err := discord.ApplicationCommandDelete(discord.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

}
