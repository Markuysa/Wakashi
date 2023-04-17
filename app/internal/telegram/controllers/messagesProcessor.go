package controllers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/services/usersService"
	"tgBotIntern/app/internal/telegram/bot"
	"tgBotIntern/app/internal/telegram/helpers"
)

type MessageProcessor interface {
	HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error
	SendMessage(msg tgbotapi.MessageConfig) error
}

type MessageHandler struct {
	Bot          *bot.Bot
	UsersService *usersService.UsersService
}

func NewMessageHandler(bot *bot.Bot, usersService *usersService.UsersService) *MessageHandler {
	return &MessageHandler{
		Bot:          bot,
		UsersService: usersService,
	}
}

func New(bot *bot.Bot) *MessageHandler {
	return &MessageHandler{Bot: bot}
}

// SendMessage sends a message to user
func (c *MessageHandler) SendMessage(msg tgbotapi.MessageConfig) error {
	_, err := c.Bot.Client.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *MessageHandler) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	case "open":
		msg.ReplyMarkup = helpers.NumericKeyboard
	case "start":
		msg.Text = "start"
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}
	err := p.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
