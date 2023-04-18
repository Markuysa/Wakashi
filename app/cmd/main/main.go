package main

import (
	"context"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/database/config"
	usersService2 "tgBotIntern/app/internal/services/usersService"
	bot2 "tgBotIntern/app/internal/telegram/bot"
	telegramConfig "tgBotIntern/app/internal/telegram/config"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/internal/telegram/worker"
)

func main() {
	ctx := context.Background()
	tgConfig, err := telegramConfig.New()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// DATABASE
	dbConfig, _ := config.New()
	tgDB := database.New(ctx, dbConfig)

	tokenDB, err := bbolt.Open("users.db", 0600, nil)

	err = tokenDB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer tokenDB.Close()

	bot, err := bot2.New(tgConfig, tokenDB)
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
