package controllers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"os"
)

type AdministratorController interface {
	HandleAddUser(msg tgbotapi.MessageConfig) error
	HandleLoginUser(msg tgbotapi.MessageConfig) error
}
type AdminHandler struct {
}

func (h *AdminHandler) HandleAddUser(msg tgbotapi.MessageConfig) error {

	requestURL := fmt.Sprintf("http://localhost:%d/register", 8080)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	return nil
}
func (h *AdminHandler) HandleLoginUser(msg tgbotapi.MessageConfig) error {

	// TODO Do request here

	return nil
}
