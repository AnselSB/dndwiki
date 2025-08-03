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
	bot.Run() // simply call for the bot to be ran
}
