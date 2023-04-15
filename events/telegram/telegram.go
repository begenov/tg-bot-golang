package telegram

import "github.com/begenov/tg-bot-golang/clients/telegramclients"

type Processor struct {
	tg     *telegramclients.Client
	offset int
	// storage
}

func NewProcessor(client *telegramclients.Client) {

}
