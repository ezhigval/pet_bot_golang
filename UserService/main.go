package main

import (
	"log"

	"pet_bot/internal/bot"
	"pet_bot/internal/config"
)

func main() {
	cfg := config.Load()

	tgBot, err := bot.New(cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	tgBot.Start()
}
