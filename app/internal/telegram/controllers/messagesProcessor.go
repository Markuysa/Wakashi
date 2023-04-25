package controllers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strings"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/metrics"
	"tgBotIntern/app/internal/service"
	"tgBotIntern/app/internal/telegram/bot"
	"tgBotIntern/app/internal/telegram/helpers"
	"tgBotIntern/app/internal/ui/messages"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"tgBotIntern/app/pkg/auth/service/usersService"
	"time"
)

// MessageHandler struct is the main handlers layer, that
// communicates with customer.
// It contains services and uses them to interact with user
// and handle the requests.
// It encapsulates the entire business logic of the bot.
type MessageHandler struct {
	tgClient           *bot.TgClientWrapper
	logger             *zap.Logger
	usersService       usersService.UsersRepositoryService
	tokenService       tokenService.TokenManager
	adminService       service.AdministratorRights
	shogunService      service.ShogunRights
	daimyoService      service.DaimyoRights
	samuraiService     service.SamuraiRights
	collectorService   service.CollectorRights
	cardService        service.CardRights
	relationService    service.RelationsServiceMethods
	transactionService service.TransactionProcessor
}

func NewMessageHandler(tgClient *bot.TgClientWrapper,
	usersService usersService.UsersRepositoryService,
	tokenRepos tokenService.TokenManager,
	adminService service.AdministratorRights,
	shogunService service.ShogunRights,
	daimyoService service.DaimyoRights,
	samuraiService service.SamuraiRights,
	collectorService service.CollectorRights,
	cardService service.CardRights,
	relationService service.RelationsServiceMethods,
	transactionService service.TransactionProcessor,
	logger *zap.Logger) *MessageHandler {
	return &MessageHandler{tgClient: tgClient,
		usersService:       usersService,
		tokenService:       tokenRepos,
		adminService:       adminService,
		shogunService:      shogunService,
		daimyoService:      daimyoService,
		samuraiService:     samuraiService,
		collectorService:   collectorService,
		cardService:        cardService,
		relationService:    relationService,
		transactionService: transactionService,
		logger:             logger,
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

func (h *MessageHandler) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	// Default commands
	case "start":
		msg.Text = messages.GreetingMessage
		msg.ReplyMarkup = helpers.GetKeyboard(-1)
		return h.SendMessage(msg)
	case "register":
		return h.handleRegister(ctx, msg, message)
	case "reset_password":
		return h.handleResetPassword(ctx, msg, message)
	case "login":
		return h.handleLogin(ctx, msg, message)
	case "exit":
		return h.handleExit(ctx, msg, message)
	case "status":
		return h.handleStatus(ctx, msg, message)

	// Administrator commands
	case "admin_createEntity":
		return h.checkRoleMiddleware(h.handleAdminCreateEntity, roles.Administrator)(ctx, msg, message)
	case "admin_createCard":
		return h.checkRoleMiddleware(h.handleAdminCreateCard, roles.Administrator)(ctx, msg, message)
	case "admin_bindSlave":
		return h.checkRoleMiddleware(h.handleAdminBindSlave, roles.Administrator)(ctx, msg, message)
	case "admin_bindCard":
		return h.checkRoleMiddleware(h.handleAdminBindCardToDaimyo, roles.Administrator)(ctx, msg, message)
	case "admin_entityData":
		return h.checkRoleMiddleware(h.handleAdminEntityData, roles.Administrator)(ctx, msg, message)

	// Shogun commands
	case "shogun_getSlavesList":
		return h.checkRoleMiddleware(h.handleShogunGetDaimyoSlavesList, roles.Shogun)(ctx, msg, message)
	case "shogun_createCard":
		return h.checkRoleMiddleware(h.handleShogunCreateCard, roles.Shogun)(ctx, msg, message)
	case "shogun_getSlavesData":
		return h.checkRoleMiddleware(h.handleShogunGetSamuraiSlaveData, roles.Shogun)(ctx, msg, message)

	// Daimyo commands
	case "daimyo_getCards":
		return h.checkRoleMiddleware(h.handleDaimyoGetCards, roles.Daimyo)(ctx, msg, message)
	case "daimyo_increase":
		return h.checkRoleMiddleware(h.handleIncCardRequest, roles.Daimyo)(ctx, msg, message)
	case "daimyo_getSamurai":
		return h.checkRoleMiddleware(h.handleDaimyoGetSamurai, roles.Daimyo)(ctx, msg, message)
	case "daimyo_getTurnover":
		return h.checkRoleMiddleware(h.handleDaimyoGetSamuraiTurnover, roles.Daimyo)(ctx, msg, message)
	case "daimyo_getCardsTotal":
		return h.checkRoleMiddleware(h.handleDaimyoGetCardsTotal, roles.Daimyo)(ctx, msg, message)
	case "daimyo_bindShogun":
		return h.checkRoleMiddleware(h.handleDaimyoBindShogun, roles.Daimyo)(ctx, msg, message)

		// Samurai commands
	case "samurai_getTurnover":
		return h.checkRoleMiddleware(h.handleSamuraiGetTurnover, roles.Samurai)(ctx, msg, message)
	case "samurai_bindDaimyo":
		return h.checkRoleMiddleware(h.handleSamuraiBindDaimyo, roles.Samurai)(ctx, msg, message)

	// Collector commands
	case "collector_performInc":
		return h.checkRoleMiddleware(h.handleCollectorIncreaseCard, roles.Collector)(ctx, msg, message)
	case "collector_showTransactions":
		return h.checkRoleMiddleware(h.handleCollectorGetTransactions, roles.Collector)(ctx, msg, message)
	default:
		return h.handleDefaultCommands(msg)
	}
}

// checkRoleMiddleware is a decorator to check user role before call the function
func (h *MessageHandler) checkRoleMiddleware(next func(context.Context, tgbotapi.MessageConfig, *tgbotapi.Message) error,
	roleID int) func(context.Context, tgbotapi.MessageConfig, *tgbotapi.Message) error {
	return func(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
		logMsg := fmt.Sprintf("Endpoint: /%s called by %s", message.Command(), message.From.UserName)
		h.logger.Info(logMsg, zap.String("Role", roles.GetRoleString(roleID)))
		isValid, err := h.usersService.IsUserSessionValid(ctx, message.From.UserName, roleID)
		if err != nil {
			msg.Text = "failed to get your status!"
			return h.SendMessage(msg)
		}
		if isValid {
			startTime := time.Now()
			err := next(ctx, msg, message)
			duration := time.Since(startTime)
			metrics.SummaryResponseTime.Observe(duration.Seconds())
			metrics.HistogramResponseTime.
				WithLabelValues("Status: OK").
				Observe(duration.Seconds())
			if err == nil {
				logMsg := fmt.Sprintf("Endpoint: /%s called by %s handled", message.Command(), message.From.UserName)
				h.logger.Info(logMsg, zap.String("Role", roles.GetRoleString(roleID)))
				return nil
			} else {
				return err
			}
		} else {
			msg.Text = "You don't have rights to call this endpoint!"
			return h.SendMessage(msg)
		}
	}
}

// handleRegister handles /register ... command
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

// handleLogin /login ... command
func (h *MessageHandler) handleLogin(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 1 {
		msg.Text = "not enough arguments in register command"
		return h.SendMessage(msg)
	}
	username := message.From.UserName
	// ignoring the error because it occurs when there is no user authorized
	userSession, _ := h.tokenService.GetUserSession(ctx, username)
	if userSession != nil {
		msg.Text = "You already authorized"
		return h.SendMessage(msg)
	}
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

// exit command
func (h *MessageHandler) handleExit(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	err := h.tokenService.ResetUserSession(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "failed to exit: " + err.Error()
		return h.SendMessage(msg)
	}
	msg.Text = "Successfully exited."
	return h.SendMessage(msg)
}

// status command
func (h *MessageHandler) handleStatus(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	_, err := h.tokenService.GetUserSession(ctx, message.From.UserName)
	if err != nil {
		msg.Text = "You should authorize first"
		return h.SendMessage(msg)
	}
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

// reset command
func (h *MessageHandler) handleResetPassword(ctx context.Context, msg tgbotapi.MessageConfig, message *tgbotapi.Message) error {
	params := strings.Split(message.Text, " ")[1:]
	if len(params) != 2 {
		msg.Text = "not enough arguments in create entity command"
		return h.SendMessage(msg)
	}
	token := strings.TrimSpace(strings.Split(params[0], "=")[1])
	newPassword := strings.TrimSpace(strings.Split(params[2], "=")[1])
	if len(token) == 0 {
		msg.Text = "Type your token!"
		return h.SendMessage(msg)
	}
	isValid, err := h.tokenService.IsTokenValid(ctx, token, message.From.UserName)
	if err != nil {
		msg.Text = err.Error()
		return h.SendMessage(msg)
	}
	if isValid {
		err := h.usersService.UpdatePassword(ctx, message.From.UserName, newPassword)
		if err != nil {
			msg.Text = "Failed to reset your password, try later."
			return h.SendMessage(msg)
		}
		msg.Text = "Successfully updated your password!"
		return h.SendMessage(msg)
	} else {
		msg.Text = "Incorrect token!"
		return h.SendMessage(msg)
	}
}

// returns text information_docs about commands
func (h *MessageHandler) handleDefaultCommands(msg tgbotapi.MessageConfig) error {
	switch msg.Text {

	case helpers.LoginButton:
		{
			msg.Text = messages.LoginMessage
		}
	case helpers.RegisterButton:
		{
			msg.Text = messages.RegisterMessage
		}
	case helpers.Info:
		{
			msg.Text = messages.InfoMessage
		}
	case helpers.Exit:
		{
			msg.Text = messages.ExitMessage
		}
	case helpers.ResetPassword:
		{
			msg.Text = messages.ResetMessage
		}
	// Admin buttons handlers
	case helpers.AdminCreateEntity:
		{
			msg.Text = messages.AdminDoc_createEntity
		}
	case helpers.AdminCreateCard:
		{
			msg.Text = messages.AdminDoc_createCard
		}
	case helpers.AdminBindSlaveToMaster:
		{
			msg.Text = messages.AdminDoc_bindSlave
		}
	case helpers.AdminVBindCardToDaimyo:
		{
			msg.Text = messages.AdminDoc_bindCardToDaimyo
		}
	case helpers.AdminGetEntityData:
		{
			msg.Text = messages.AdminDoc_getEntityData
		}
		// Shogun buttons handlers
	case helpers.ShogunGetSlavesList:
		{
			msg.Text = messages.ShodunDoc_getSlavesList
		}
	case helpers.ShogunCreateCard:
		{
			msg.Text = messages.ShodunDoc_createCard
		}
	case helpers.ShogunGetSlaveData:
		{
			msg.Text = messages.ShodunDoc_getSlavesData
		}
		// Daimyo buttons handlers
	case helpers.DaimyoGetCards:
		{
			msg.Text = messages.DaimyoDoc_getCards
		}
	case helpers.DaimyoCreateCardRequest:
		{
			msg.Text = messages.DaimyoDoc_increase
		}
	case helpers.DaimyoGetSamuraiList:
		{
			msg.Text = messages.DaimyoDoc_getSamurai
		}
	case helpers.DaimyoGetSamuraiTurnover:
		{
			msg.Text = messages.DaimyoDoc_getTurnover
		}
	case helpers.DaimyoGetCardsTotal:
		{
			msg.Text = messages.DaimyoDoc_getCardsTotal
		}
	case helpers.DaimyoBindShogun:
		{
			msg.Text = messages.DaimyoDoc_bindShogun
		}
		// Samurai buttons handlers
	case helpers.SamuraiGetTurnover:
		{
			msg.Text = messages.SamuraiDoc_getTurnover
		}

	case helpers.SamuraiBindDaimyo:
		{
			msg.Text = messages.SamuraiDoc_bindDaimyo
		}

	case helpers.CollectorProccessRequest:
		{
			msg.Text = messages.CollectorDoc_performInc
		}

	case helpers.CollectorShowTransactions:
		{
			msg.Text = messages.CollectorDoc_showTranasctions
		}
	}
	return h.SendMessage(msg)
}
