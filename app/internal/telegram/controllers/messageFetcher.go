package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/telegram/bot"
)

// MessageFetcher interface represents the contract
// of object that will fetch the messages from tg bot
type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Stop()
}

type FetcherWorker struct {
	Bot *bot.Bot
}

func NewFetcherWorker(bot *bot.Bot) *FetcherWorker {
	return &FetcherWorker{Bot: bot}
}

// Start method returns updates channel of the bot
func (c *FetcherWorker) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.Bot.Client.GetUpdatesChan(u)
}

// Stop method stops receiving updates from the bot
func (c *FetcherWorker) Stop() {
	c.Bot.Client.StopReceivingUpdates()
}
