package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/telegram/bot"
)

type FetcherWorker struct {
	tgClient *bot.TgClientWrapper
}

func NewFetcherWorker(bot *bot.TgClientWrapper) *FetcherWorker {
	return &FetcherWorker{tgClient: bot}
}

// Start method returns updates channel of the bot
func (c *FetcherWorker) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.tgClient.Client.GetUpdatesChan(u)
}

// Stop method stops receiving updates from the bot
func (c *FetcherWorker) Stop() {
	c.tgClient.Client.StopReceivingUpdates()
}
