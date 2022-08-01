package main

import (
	"discord-bot/config"
	discord_bot "discord-bot/internal/discord-bot"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("read arg error")
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("read config error: %s", err)
	}
	var configStruct config.Main
	err = yaml.Unmarshal(file, configStruct)
	if err != nil {
		log.Fatalf("unmarshal config file error: %s", err)
	}
	err = discord_bot.Run(configStruct)
	if err != nil {
		log.Fatal(err)
	}
}
