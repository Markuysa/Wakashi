package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Controller interface {
	AdministratorController
}

type Controllers struct {
	client tgbotapi.BotAPI
}
