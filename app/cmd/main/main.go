package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/hedzr/errors.v3"
	"log"
	"os"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/database/config"
	logger2 "tgBotIntern/app/internal/logger"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("tg_config_path"))
	tgConfig, err := telegramConfig.New()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logger2.InitLogger()
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
	daimyoService := service.NewDaimyoService(cardsService, userService, relationService, transactionService)
	samuraiService := service.NewSamuraiService(relationService, cardsService, userService)
	collectorService := service.NewCollectorService(cardsService, transactionService)

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
		logger,
	)

	// WORKERS
	msgListenerWorker := worker.NewMessageListenerWorker(msgListener, msgHandler)

	// Start auth server
	authServer := server.NewAuthSerer(AuthServerPort, sessionsDB, userService, logger)
	go func() {
		err := authServer.Start()
		if err != nil {
			log.Fatal(errors.New("failed to start authServer:%v", err))
		}
	}()

	// Start receiving messages
	logger.Info("Started telegram bot")

	err = msgListenerWorker.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
