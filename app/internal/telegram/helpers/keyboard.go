package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/constants/roles"
)

var (
	adminKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("create entity"),
			tgbotapi.NewKeyboardButton("create card"),
			tgbotapi.NewKeyboardButton("bind card to daimyo"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("bind slave to master"),
			tgbotapi.NewKeyboardButton("get entity data"),
		),
	)
)

func GetKeyboard(roleID int) tgbotapi.ReplyKeyboardMarkup {
	switch roleID {
	case roles.Administrator:
		{
			return adminKeyboard
		}
	default:
		// return default keyboard
		return adminKeyboard
	}
}
