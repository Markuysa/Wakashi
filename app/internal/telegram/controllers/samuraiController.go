package controllers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

import "context"

func (h *MessageHandler) handleSamuraiGetTurnover(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	turnover, err := h.samuraiService.GetTurnover(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "No transactions yet, turnover is 0"
		return h.SendMessage(msg)
	}
	msg.Text = fmt.Sprintf("Your turnover is:%v", turnover)
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleSamuraiBindDaimyo(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	daimyoUsername := strings.TrimSpace(strings.Split(params[0], "=")[1])
	err := h.samuraiService.BindToDamiyo(ctx, daimyoUsername, message.From.UserName)
	if err != nil {
		msg.Text = "Cannot bind to Daimyo: " + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully bound!"
	return h.SendMessage(msg)
}
