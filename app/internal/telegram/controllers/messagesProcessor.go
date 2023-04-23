package controllers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/telegram/bot"
	"tgBotIntern/app/internal/telegram/helpers"
	"tgBotIntern/app/internal/ui/messages"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type MessageHandler struct {
	tgClient     *bot.TgClientWrapper
	usersService usersService.UsersRepositoryService
	tokenRepos   tokenService.TokenManager
}

func NewMessageHandler(tgClient *bot.TgClientWrapper, usersService usersService.UsersRepositoryService, tokenRepos tokenService.TokenManager) *MessageHandler {
	return &MessageHandler{tgClient: tgClient, usersService: usersService, tokenRepos: tokenRepos}
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

func (h *MessageHandler) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	// Default commands
	case "start":
		msg.Text = messages.GreetingMessage
		return h.SendMessage(msg)
	case "register":
		return h.handleRegister(ctx, msg, message)
	case "login":
		return h.handleLogin(ctx, msg, message)
	case "exit":
		return h.handleExit(ctx, msg, message)
	case "status":
		return h.handleStatus(ctx, msg, message)

	// Administrator commands
	case "admin_createEntity":
		return h.handleAdminCreateEntity(ctx, msg, message)
	case "admin_createCard":
		return h.handleAdminCreateCard(ctx, msg, message)
	case "admin_bindDaimyo":
		return h.handleAdminBindDaimyo(ctx, msg, message)
	case "admin_bindSlave":
		return h.handleAdminVindSlave(ctx, msg, message)
	case "admin_entityData":
		return h.handleAdminEntityData(ctx, msg, message)

	// Shogun commands
	case "shogun_getSlavesList":
		return h.handleShogunGetSlavesList(ctx, msg, message)
	case "shogun_createCard":
		return h.handleShogunCreateCard(ctx, msg, message)
	case "shogun_bindCardToSamurai":
		return h.handleShogunBindCardToSamurai(ctx, msg, message)
	case "shogun_getSlavesData":
		return h.handleShogunGetSlavesData(ctx, msg, message)

	// Daimyo commands
	case "daimyo_getCards":
		return h.handleDaimyoGetCards(ctx, msg, message)
	case "daimyo_increase":
		return h.handleIncCardRequest(ctx, msg, message)
	case "daimyo_getSamurai":
		return h.handleDaimyoGetSamurai(ctx, msg, message)
	case "daimyo_getTotal":
		return h.handleDaimyoGetTotal(ctx, msg, message)
	case "daimyo_bindShogun":
		return h.handleDaimyoBindShogun(ctx, msg, message)
	// Samurai commands
	case "samurai_getTurnover":
		return h.handleSamuraiGetTurnover(ctx, msg, message)
	case "samurai_bindDaimyo":
		return h.handleSamuraiBindDaimyo(ctx, msg, message)
	// Collector commands
	case "collector_performInc":
		return h.handleCollectorIncreaseCard(ctx, msg, message)
	default:
		return h.handleDefaultCommands(ctx, msg, message)
	}
}

func (h *MessageHandler) handleRegister(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 2 {
		msg.Text = "not enough arguments in register command"
		return h.SendMessage(msg)
	}
	username := message.From.UserName
	password := strings.TrimSpace(strings.Split(params[0], "=")[1])
	role := strings.TrimSpace(strings.Split(params[1], "=")[1])
	roleID := roles.GetRoleID(role)
	err := h.usersService.RegisterUser(ctx, username, password, roleID)
	if err != nil {
		msg.Text = "unable to register user:" + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully registered!"
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleLogin(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in register command"
		return h.SendMessage(msg)
	}
	username := message.From.UserName
	password := strings.TrimSpace(strings.Split(params[0], "=")[1])
	role, _ := h.usersService.GetRoleID(ctx, username)
	tokens, err := h.usersService.AuthorizeUser(ctx, username, password)
	if err != nil {
		msg.Text = "failed to authorize:" + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully authorized! Your access token:" + tokens.AccessToken
	msg.ReplyMarkup = helpers.GetKeyboard(role)
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleExit(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	err := h.tokenRepos.ResetUserSession(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "failed to exit: " + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully exited."
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleStatus(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roles.Administrator)
	if err != nil {
		msg.Text = "failed to get your status!"
		return h.SendMessage(msg)
	}
	if isValid {
		user, err := h.usersService.GetUser(ctx, message.From.UserName)
		if err != nil {
			msg.Text = "failed to get your data!"
			return h.SendMessage(msg)
		}
		msg.Text = fmt.Sprintf(`
			Username: %v, 
			Role: %v,
		`, user.Username, roles.GetRoleString(user.Role))
		return h.SendMessage(msg)
	}
	msg.Text = "You don't have enough rights to use that command!"
	return h.SendMessage(msg)
}
func (h *MessageHandler) handleDefaultCommands(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	switch msg.Text {
	case "create entity":
		{
			msg.Text = messages.AdminDoc_createEntity
			return h.SendMessage(msg)
		}
	case "create card":
		{
			msg.Text = messages.AdminDoc_createCard
			return h.SendMessage(msg)
		}
	}
	// TODO finish cases
	return h.SendMessage(msg)
}
