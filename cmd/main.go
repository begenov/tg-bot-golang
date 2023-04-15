package main

import (
	"flag"
	"log"

	"github.com/begenov/tg-bot-golang/clients/telegramclients"
	eventconsumer "github.com/begenov/tg-bot-golang/consumer/event-consumer"
	"github.com/begenov/tg-bot-golang/events/telegram"
	"github.com/begenov/tg-bot-golang/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	tgClient := telegramclients.NewClient(mustToken(), tgBotHost)

	eventsProcessor := telegram.NewProcessor(&tgClient, files.NewStorage(storagePath))

	log.Print("service started")

	consumer := eventconsumer.NewConsumer(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
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
