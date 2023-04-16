package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/hedzr/errors.v3"
	"log"
	"tgBotIntern/app/intenal/telegram/config"
)

// Client is a struct of telegram bot api client
type Client struct {
	client *tgbotapi.BotAPI
}

// New creates new tgBot client
func New(tgConfig *config.Config) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tgConfig.ApiToken)

	if err != nil {
		return nil, errors.New("failed to init tg bot api:%v", err)
	}
	return &Client{
		client: client,
	}, nil
}

// Start method returns updates channel of the bot
func (c *Client) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.client.GetUpdatesChan(u)
}

// Stop method stops receiving updates from the bot
func (c *Client) Stop() {
	c.client.StopReceivingUpdates()
}

// SendMessage sends a message to user
func (c *Client) SendMessage(msg tgbotapi.MessageConfig) {
	_, err := c.client.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
