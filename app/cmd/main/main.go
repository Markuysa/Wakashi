package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/hedzr/errors.v3"
	"log"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/database/config"
	"tgBotIntern/app/internal/telegram/bot"
	telegramConfig "tgBotIntern/app/internal/telegram/config"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/internal/telegram/worker"
	"tgBotIntern/app/pkg/auth/server"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	usersService2 "tgBotIntern/app/pkg/auth/service/usersService"
	"tgBotIntern/app/pkg/auth/tokenDb"
	"time"
)

const (
	AuthServerPort = ":8080"
)

func main() {
	ctx := context.Background()
	tgConfig, err := telegramConfig.New()
	if err != nil {
		log.Fatal(err)
	}

	// DATABASE - ENTITIES
	dbConfig, _ := config.New()
	tgDB := database.New(ctx, dbConfig)

	redisCLient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "islam20011",
		DB: 0,
	})
	tokenDB := tokenDb.NewTokenRepository(redisCLient, 24*time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	botClient, err := bot.New(tgConfig)

	// SERVICES - USE CASE
	tokensService := tokenService.NewTokenService(tokenDB)
	usersService := usersService2.NewUsersService(tgDB, tokensService)

	// CONTROLLERS - HANDLERS
	msgListener := controllers.NewFetcherWorker(botClient)
	msgHandler := controllers.NewMessageHandler(botClient, usersService, tokensService)

	// WORKERS
	msgListenerWorker := worker.NewMessageListenerWorker(msgListener, msgHandler)

	// START THE AUTHENTICATION SERVER
	authServer := server.NewAuthSerer(AuthServerPort, tokenDB, usersService)
	go func() {
		err := authServer.Start()
		if err != nil {
			log.Fatal(errors.New("failed to start authServer:%v", err))
		}
	}()

	// Start receiving messages
	log.Println("started telegram bot")
	msgListenerWorker.Run(ctx)

}
