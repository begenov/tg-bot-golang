package main

import (
	"flag"
	"log"

	tgClient "github.com/begenov/tg-bot-golang/clients/telegram"
	eventconsumer "github.com/begenov/tg-bot-golang/consumer/event-consumer"
	"github.com/begenov/tg-bot-golang/events/telegram"
	"github.com/begenov/tg-bot-golang/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoppet", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
