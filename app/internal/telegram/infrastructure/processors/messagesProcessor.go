package processors

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBotIntern/app/internal/telegram/bot"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type MessageProcessor interface {
	HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error
	SendMessage(msg tgbotapi.MessageConfig) error
	controllers.AdministratorController
}

type MessageHandler struct {
	Bot          *bot.Bot
	UsersService *usersService.UsersService
}

func (h *MessageHandler) HandleAddUser(msg tgbotapi.MessageConfig) error {
	//TODO implement me
	panic("implement me")
}

func (h *MessageHandler) HandleLoginUser(msg tgbotapi.MessageConfig) error {
	//TODO implement me
	panic("implement me")
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

type user struct {
	username string
	password string
	roleID   int
}

// SendMessage sends a message to user
func (h *MessageHandler) SendMessage(msg tgbotapi.MessageConfig) error {
	msg.ParseMode = "HTML"
	_, err := h.Bot.Client.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) HandleIncomingMessage(ctx context.Context, message *tgbotapi.Message) error {

	//get, err := h.Bot.TokenRepos.Get(message.Chat.ID, "sessions")
	//if err != nil {
	//	return errors.New("failed to get the session information")
	//}
	//fmt.Println(get)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	case "register":
		msg.Text = "<a href='http://127.0.0.1:8080/register'>Register in system</a>"
	case "login":
		return h.Bot.Controllers.HandleLoginUser(msg)
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}
	err := h.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) formUser() {

}
