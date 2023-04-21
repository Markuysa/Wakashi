package controllers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller interface {
	AdministratorController
	MessageFetcher
	MessageProcessor
}

type Controllers struct {
	client tgbotapi.BotAPI
}

type AdministratorController interface {
	HandleAddUser(msg tgbotapi.MessageConfig) error
	HandleLoginUser(msg tgbotapi.MessageConfig) error
}

// InfrastructureController handles users interactions
type InfrastructureController interface {
}

// MessageProcessor interface represents the contract
// of object that will fetch the messages from tg bot
type MessageProcessor interface {
	HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message)
	SendMessage(msg tgbotapi.MessageConfig) error
}

// MessageFetcher interface represents the contract
// of object that will fetch the messages from tg bot
type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Stop()
}
