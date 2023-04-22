package bot

import (
	// telebot /:
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/telegram/config"
)

// TgClientWrapper is a struct of telegram bot api client
type TgClientWrapper struct {
	Client *tgbotapi.BotAPI
}

// New creates new tgBot client
func New(tgConfig *config.Config) (*TgClientWrapper, error) {
	client, err := tgbotapi.NewBotAPI(tgConfig.ApiToken)

	if err != nil {
		return nil, errors.New("failed to init tg bot api:%v", err)
	}
	return &TgClientWrapper{
		Client: client,
	}, nil
}
