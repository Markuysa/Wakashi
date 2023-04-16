package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/intenal/telegram/client"
	"tgBotIntern/app/intenal/telegram/helpers"
)

type MessageProcessor struct {
	Client *client.Client
}

func New(client *client.Client) *MessageProcessor {
	return &MessageProcessor{Client: client}
}

func (p *MessageProcessor) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	case "open":
		msg.ReplyMarkup = helpers.NumericKeyboard
	case "help":
		msg.Text = "I understand /sayhi and /status."
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}
	p.Client.SendMessage(msg)
}
