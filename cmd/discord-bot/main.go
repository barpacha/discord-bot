package main

import (
	"discord-bot/config"
	"discord-bot/internal/app/discord-bot"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("read arg error")
	}
	configStruct := config.Yaml(os.Args[1])
	err := discord_bot.Run(configStruct)
	if err != nil {
		log.Fatal(err)
	}
}
