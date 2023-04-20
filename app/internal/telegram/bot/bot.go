package bot

import (
	// telebot /:
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/telegram/config"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/pkg/auth/tokenDb"
)

// Bot is a struct of telegram bot api client
type Bot struct {
	Client      *tgbotapi.BotAPI
	TokenRepos  tokenDb.TokenRepos
	Controllers controllers.Controller
}

// New creates new tgBot
func New(tgConfig *config.Config, tokenRepos tokenDb.TokenRepos, controller controllers.Controller) (*Bot, error) {
	client, err := tgbotapi.NewBotAPI(tgConfig.ApiToken)

	if err != nil {
		return nil, errors.New("failed to init tg bot api:%v", err)
	}
	return &Bot{
		Client:      client,
		TokenRepos:  tokenRepos,
		Controllers: controller,
	}, nil
}
