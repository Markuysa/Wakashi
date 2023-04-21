package controllers

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tgBotIntern/app/internal/telegram/bot"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type MessageHandler struct {
	tgClient     *bot.TgClientWrapper
	usersService usersService.UsersRepositoryService
}

func NewMessageHandler(bot *bot.TgClientWrapper, usersService usersService.UsersRepositoryService) *MessageHandler {
	return &MessageHandler{
		tgClient:     bot,
		usersService: usersService,
	}
}

// SendMessage sends a message to user
func (h *MessageHandler) SendMessage(msg tgbotapi.MessageConfig) error {
	msg.ParseMode = "HTML"
	_, err := h.tgClient.Client.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	case "register":
		params := strings.Split(message.Text, " ")[1:]
		err := h.handleRegister(ctx, msg, params)
		if err != nil {
			msg.Text = err.Error()
		}
	case "login":
		params := strings.Split(message.Text, " ")[1:]
		err := h.handleLogin(ctx, msg, params)
		if err != nil {
			msg.Text = err.Error()
		}

	case "sayhi":
		cond, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, "Administrator")
		if cond {
			msg.Text = "I'm ok."
		} else {
			msg.Text = err.Error()
		}
		h.SendMessage(msg)
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}
}

func (h *MessageHandler) handleRegister(ctx context.Context, msg tgbotapi.MessageConfig, params []string) error {
	if len(params) != 3 {
		return errors.New("not enough arguments in register command")
	}
	username := strings.TrimSpace(strings.Split(params[0], "=")[1])
	password := strings.TrimSpace(strings.Split(params[1], "=")[1])
	role := strings.TrimSpace(strings.Split(params[2], "=")[1])
	err := h.usersService.RegisterUser(ctx, username, password, role)
	if err != nil {

		return err
	}
	msg.Text = "Successfully registered!"
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleLogin(ctx context.Context, msg tgbotapi.MessageConfig, params []string) error {
	if len(params) != 2 {
		return errors.New("not enough arguments in register command")
	}
	username := strings.TrimSpace(strings.Split(params[0], "=")[1])
	password := strings.TrimSpace(strings.Split(params[1], "=")[1])
	tokens, err := h.usersService.AuthorizeUser(ctx, username, password)
	if err != nil {
		return err
	}
	msg.Text = "Successfully authorized! Your access token:" + tokens.AccessToken
	return h.SendMessage(msg)
}
