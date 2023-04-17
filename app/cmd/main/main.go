package main

import (
	"context"
	"log"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/database/config"
	usersService2 "tgBotIntern/app/internal/services/usersService"
	"tgBotIntern/app/internal/telegram/bot"
	telegramConfig "tgBotIntern/app/internal/telegram/config"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/internal/telegram/worker"
)

func main() {
	ctx := context.Background()
	// It's better to use sessions than context injection
	// TODO - replace withValue with session of currentUser
	ctx = context.WithValue(ctx, "ROLE", -1)
	tgConfig, err := telegramConfig.New()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	bot, err := bot.New(tgConfig)
	if err != nil {
		log.Fatal(err)
	}

	// DATABASE
	dbConfig, _ := config.New()
	tgDB := database.New(ctx, dbConfig)

	// SERVICES
	usersService := usersService2.New(tgDB)

	// CONTROLLERS
	msgListener := controllers.NewFetcherWorker(bot)
	msgHandler := controllers.NewMessageHandler(bot, usersService)

	// WORKERS
	msgListenerWorker := worker.NewMessageListenerWorker(msgListener, msgHandler)

	// Start receiving messages
	log.Println("started telegram bot")
	msgListenerWorker.Run(ctx)

}
