package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"tgBotIntern/app/internal/telegram/helpers"
)

import "context"

// Collector role requests handlers
func (h *MessageHandler) handleCollectorIncreaseCard(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	requestID := strings.TrimSpace(strings.Split(params[0], "=")[1])
	if len(requestID) == 0 {
		msg.Text = "Length of each parameter should be > 0!"
		return h.SendMessage(msg)
	}
	requestIDInt, err := strconv.Atoi(requestID)
	if err != nil {
		msg.Text = "Check your requestID! It should be a number without literals"
		return h.SendMessage(msg)
	}
	err = h.collectorService.HandleDaimyoIncreasementRequest(ctx, requestIDInt)
	if err != nil {
		msg.Text = "Unable to handle request: " + requestID
		return h.SendMessage(msg)
	}
	msg.Text = "The card balance successfully changed"
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleCollectorGetTransactions(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	transactions, err := h.transactionService.GetUnhandledTransactions(ctx)
	if err != nil {
		msg.Text = "Unable to get transactions!"
		return h.SendMessage(msg)
	}
	msg.Text = helpers.FormTransactions(transactions)
	return h.SendMessage(msg)
}
