package main

import (
	"flag"
	"log"

	"github.com/begenov/tg-bot-golang/clients/telegramclients"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegramclients.NewClient(mustToken(), tgBotHost)
}

func mustToken() string {
	token := flag.String(
		"token-bot",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
