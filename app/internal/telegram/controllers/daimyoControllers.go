package controllers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"tgBotIntern/app/internal/telegram/helpers"
)
import "context"

// Daimyo role requests handlers

func (h *MessageHandler) handleDaimyoGetCards(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	cards, err := h.daimyoService.GetCardsList(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "Failed to get cards" + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = helpers.FormCardsList(cards)
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleIncCardRequest(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 2 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	cardNumber := strings.TrimSpace(strings.Split(params[0], "=")[1])
	incValue := strings.TrimSpace(strings.Split(params[1], "=")[1])
	cardNumberInt, err := strconv.Atoi(cardNumber)
	if err != nil || len(cardNumber) != 16 {
		msg.Text = "Check your cardNumber! It should be 16-digit number without spaces in between"
		return h.SendMessage(msg)
	}
	incValueFloat, err := strconv.ParseFloat(incValue, 64)
	if err != nil {
		msg.Text = "Check your incrementalValue"
		return h.SendMessage(msg)
	}
	err = h.daimyoService.CreateCardIncreasementRequest(ctx, cardNumberInt, incValueFloat, message.From.UserName)
	if err != nil {
		msg.Text = "Failed to create request"
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully created!"
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleDaimyoGetSamurai(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	samurai, err := h.daimyoService.GetSamuraiList(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "Cannot get your samurai list: " + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = helpers.FormSlavesList(samurai)
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleDaimyoGetSamuraiTurnover(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	samuraiUsername := strings.TrimSpace(strings.Split(params[0], "=")[1])
	turnover, err := h.daimyoService.GetSamuraiTurnover(ctx, message.From.UserName, samuraiUsername)
	if err != nil {
		msg.Text = err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = fmt.Sprintf("The turnover of %s is %v", samuraiUsername, turnover)
	return h.SendMessage(msg)

}

func (h *MessageHandler) handleDaimyoGetCardsTotal(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {

	total, err := h.daimyoService.GetCardsTotal(ctx, message.From.UserName)
	if err != nil {
		msg.Text = err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = fmt.Sprintf("The tot	al is:%v", total)
	return h.SendMessage(msg)

}

func (h *MessageHandler) handleDaimyoBindShogun(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	shogunUsername := strings.TrimSpace(strings.Split(params[0], "=")[1])
	err := h.daimyoService.BindShogun(ctx, shogunUsername, message.From.UserName)
	if err != nil {
		msg.Text = "Cannot bind to shogun"
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully binded to shogun: " + shogunUsername
	return h.SendMessage(msg)
}
