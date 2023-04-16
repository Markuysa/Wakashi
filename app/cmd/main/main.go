package main

import (
	"context"
	"log"
	"tgBotIntern/app/intenal/telegram/client"
	telegramConfig "tgBotIntern/app/intenal/telegram/config"
	"tgBotIntern/app/intenal/telegram/handlers"
	"tgBotIntern/app/intenal/worker"
)

func main() {
	ctx := context.Background()
	tgConfig, err := telegramConfig.New()
	if err != nil {
		log.Fatal(err)
	}
	tgClient, err := client.New(tgConfig)
	if err != nil {
		log.Fatal(err)
	}
	tgProcessor := handlers.New(tgClient)

	msgListener := worker.NewMessageListenerWorker(tgClient, tgProcessor)

	// Start receiving messages

	log.Println("started telegram bot")
	msgListener.Run(ctx)
}
