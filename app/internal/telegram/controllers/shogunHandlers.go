package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/internal/helpers/encoder"
	"tgBotIntern/app/internal/telegram/helpers"
)
import "context"

func (h *MessageHandler) handleShogunGetDaimyoSlavesList(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	slaves, err := h.shogunService.GetSlavesList(ctx, message.From.UserName, roles.Daimyo)
	if err != nil {
		return err
	}
	slavesList := helpers.FormSlavesList(slaves)
	if len(slavesList) == 0 {
		msg.Text = "You don't have any slaves"
	} else {
		msg.Text = slavesList
	}
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleShogunCreateCard(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 4 {
		msg.Text = "not enough arguments in create card command"
		return h.SendMessage(msg)
	}
	number := strings.TrimSpace(strings.Split(params[0], "=")[1])
	bank := strings.TrimSpace(strings.Split(params[1], "=")[1])
	owner := strings.TrimSpace(strings.Split(params[2], "=")[1])
	cvv := strings.TrimSpace(strings.Split(params[3], "=")[1])
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		msg.Text = "The card number should contain only digits without spaces in between!"
		return h.SendMessage(msg)
	}
	bankID, err := strconv.Atoi(bank)
	if err != nil {
		msg.Text = "The bank ID should contain only digits without spaces in between!"
		return h.SendMessage(msg)
	}
	ownerID, err := h.usersService.GetUserID(ctx, owner)
	if err != nil {
		msg.Text = "No user with that username"
		return h.SendMessage(msg)
	}
	cvvHash, err := encoder.Encode(cvv)
	if err != nil {
		msg.Text = "Something wrong with your cvv"
		return h.SendMessage(msg)
	}
	err = h.shogunService.CreateCard(ctx, entity.Card{
		DaimyoID:     ownerID,
		IssuerBankID: bankID,
		CvvCode:      cvvHash,
		CardNumber:   int64(numberInt),
		DailyLimit:   2_000_000,
		Total:        0,
	})
	if err != nil {
		msg.Text = "unable to craete card: " + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully created new card with number:" + number + "!"
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleShogunGetSamuraiSlaveData(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in create card command"
		return h.SendMessage(msg)
	}
	slaveUsername := strings.TrimSpace(strings.Split(params[0], "=")[1])

	slave, err := h.usersService.GetUser(ctx, slaveUsername)
	if err != nil {
		msg.Text = "User with username: " + slave.Username + " not found"
		return h.SendMessage(msg)
	}
	turnover, err := h.transactionService.GetTurnover(ctx, slave.Username)
	s := strconv.FormatFloat(turnover, 'E', -1, 64)
	if err != nil {
		msg.Text = "Cannot calculate turnover of Samurai: " + slaveUsername
		return h.SendMessage(msg)
	}
	msg.Text = helpers.FormUser(slave) + "\n" + s
	return h.SendMessage(msg)
}
