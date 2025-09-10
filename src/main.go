package main

import (
	"fmt"
	"os"
	bot "wikiBot/src/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR LOADING ENV")
		return
	}
	bot.BotToken = os.Getenv("DISCORD_BOT_TOKEN")
	bot.DumpChannelID = os.Getenv("DUMP_CHANNEL_ID")
	bot.SpellFormURL = os.Getenv("SPELL_FORM_URL")
	bot.ItemFormURL = os.Getenv("ITEM_FORM_URL")
	bot.Run() // simply call for the bot to be ran
}
