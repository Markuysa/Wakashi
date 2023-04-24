package controllers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/internal/helpers/encoder"
	"tgBotIntern/app/internal/telegram/helpers"
)

// TODO create middleware function

func (h *MessageHandler) handleAdminCreateEntity(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
		params := strings.Split(message.Text, " ")[1:]
		if len(params) != 3 {
			msg.Text = "not enough arguments in create entity command"
			return h.SendMessage(msg)
		}
		username := strings.TrimSpace(strings.Split(params[0], "=")[1])
		password := strings.TrimSpace(strings.Split(params[1], "=")[1])
		role := strings.TrimSpace(strings.Split(params[2], "=")[1])
		roleID := roles.GetRoleID(role)
		err := h.usersService.RegisterUser(ctx, username, password, roleID)
		if err != nil {
			msg.Text = "unable to register user:" + err.Error()
			return h.SendMessage(msg)
		}
		msg.Text = "Successfully added new entity with Role:" + role + "!"
	} else {
		msg.Text = "You don't have rights to call this endpoint!"
	}
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleAdminCreateCard(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
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
			return err
		}
		cvvHash, err := encoder.Encode(cvv)
		if err != nil {
			msg.Text = "Something wrong with your cvv"
			return h.SendMessage(msg)
		}
		err = h.adminService.CreateCard(ctx, entity.Card{
			DaimyoID:     ownerID,
			IssuerBankID: bankID,
			CvvCode:      cvvHash,
			CardNumber:   int64(numberInt),
			DailyLimit:   2_000_000,
			Total:        0,
		})
		if err != nil {
			msg.Text = "unable to register user:" + err.Error()
			return h.SendMessage(msg)
		}
		msg.Text = "Successfully created new card with number:" + number + "!"
	} else {
		msg.Text = "You don't have rights to call this endpoint!"
	}
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleAdminBindSlave(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
		params := strings.Split(message.Text, " ")[1:]
		if len(params) != 2 {
			msg.Text = "not enough arguments in bind slave command"
			return h.SendMessage(msg)
		}
		master := strings.TrimSpace(strings.Split(params[0], "=")[1])
		slave := strings.TrimSpace(strings.Split(params[1], "=")[1])
		getMaster, err := h.usersService.GetUser(ctx, master)
		if err != nil {
			msg.Text = "User: " + master + " not found"
			return h.SendMessage(msg)
		}
		getSlave, err := h.usersService.GetUser(ctx, slave)
		if err != nil {
			msg.Text = "User: " + slave + " not found"
			return h.SendMessage(msg)
		}
		if getMaster.Role >= getSlave.Role {
			msg.Text = "Role of user: " + getMaster.Username + " is less than or equal to " + getSlave.Username + "'s role"
		}
		err = h.relationService.Bind(ctx, master, slave)
		if err != nil {
			msg.Text = "Error of binding the users"
			return h.SendMessage(msg)
		}
		msg.Text = "Successfully bound!"
	} else {
		msg.Text = "You don't have rights to call this endpoint!"
	}
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleAdminBindCardToDaimyo(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
		params := strings.Split(message.Text, " ")[1:]
		if len(params) != 2 {
			msg.Text = "not enough arguments in bind card command"
			return h.SendMessage(msg)
		}
		cardNumber := strings.TrimSpace(strings.Split(params[0], "=")[1])
		daimyoUsername := strings.TrimSpace(strings.Split(params[1], "=")[1])
		err := h.cardService.BindToDaimyo(ctx, cardNumber, daimyoUsername)
		if err != nil {
			msg.Text = "Failed to bind your card to daimyo!"
			return h.SendMessage(msg)
		}
	} else {
		msg.Text = "You don't have rights to call this endpoint!"
	}
	return h.SendMessage(msg)
}

func (h *MessageHandler) handleAdminEntityData(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
		params := strings.Split(message.Text, " ")[1:]
		if len(params) != 1 {
			msg.Text = "not enough arguments in get entity data command"
			return h.SendMessage(msg)
		}
		username := strings.TrimSpace(strings.Split(params[0], "=")[1])
		report, err := h.adminService.GetEntityReport(ctx, username)
		if err != nil {
			msg.Text = "Failed to create report"
			return h.SendMessage(msg)
		}
		msg.Text = helpers.FormMessage(report)
	} else {
		msg.Text = "You don't have rights to call this endpoint!"
	}
	return h.SendMessage(msg)
}
