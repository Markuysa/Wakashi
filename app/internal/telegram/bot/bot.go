package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/telegram/config"
)

// Client is a struct of telegram bot api client
type Bot struct {
	Client *tgbotapi.BotAPI
}

// New creates new tgBot
func New(tgConfig *config.Config) (*Bot, error) {
	client, err := tgbotapi.NewBotAPI(tgConfig.ApiToken)

	if err != nil {
		return nil, errors.New("failed to init tg bot api:%v", err)
	}
	return &Bot{
		Client: client,
	}, nil
}
