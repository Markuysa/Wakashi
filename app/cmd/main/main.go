package main

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"log"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/database/config"
	"tgBotIntern/app/internal/service"
	"tgBotIntern/app/internal/telegram/bot"
	telegramConfig "tgBotIntern/app/internal/telegram/config"
	"tgBotIntern/app/internal/telegram/controllers"
	"tgBotIntern/app/internal/telegram/worker"
	configSession "tgBotIntern/app/pkg/auth/config"
	sessionDB "tgBotIntern/app/pkg/auth/database"
	"tgBotIntern/app/pkg/auth/server"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	usersService "tgBotIntern/app/pkg/auth/service/usersService"
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
	dbConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	sessionDBCOnfig, err := configSession.New()
	if err != nil {
		log.Fatal(err)
	}
	tgConn := database.NewDBConnection(ctx, dbConfig)

	usersDB := database.NewUsersDB(tgConn)
	cardDB := database.NewCardsDB(tgConn)
	relationDB := database.NewRelationDB(tgConn)
	transactionsDB := database.NewTransactionRepository(tgConn)
	sessionsDB := sessionDB.NewTokenRepository(sessionDBCOnfig)

	if err != nil {
		log.Fatal(err)
	}

	botClient, err := bot.New(tgConfig)

	// SERVICES - USE CASE
	tokensService := tokenService.NewTokenService(sessionsDB)
	userService := usersService.NewUsersService(usersDB, tokensService, 24*time.Hour, time.Hour)
	cardsService := service.NewCardService(cardDB, userService)
	relationService := service.NewRelationsService(relationDB, userService)
	transactionService := service.NewTransactionsService(transactionsDB, userService)
	adminService := service.NewAdministratorService(userService, cardsService, relationService, transactionService)
	shogunService := service.NewShogunService(userService, cardsService)
	daimyoService := service.NewDaimyoService(cardsService, userService, relationService)
	samuraiService := service.NewSamuraiService(relationService, cardsService)
	collectorService := service.NewCollectorService(cardsService)

	// CONTROLLERS - HANDLERS
	msgListener := controllers.NewFetcherWorker(botClient)
	msgHandler := controllers.NewMessageHandler(
		botClient,
		userService,
		tokensService,
		adminService,
		shogunService,
		daimyoService,
		samuraiService,
		collectorService,
		cardsService,
		relationService,
		transactionService,
	)

	// WORKERS
	msgListenerWorker := worker.NewMessageListenerWorker(msgListener, msgHandler)

	// START THE AUTHENTICATION SERVER
	authServer := server.NewAuthSerer(AuthServerPort, sessionsDB, userService)
	go func() {
		err := authServer.Start()
		if err != nil {
			log.Fatal(errors.New("failed to start authServer:%v", err))
		}
	}()

	// Start receiving messages
	log.Println("started telegram bot")
	err = msgListenerWorker.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
