package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"tgBotIntern/app/internal/database/config"
)

type TelegramDB interface {
	UsersDB
	CardsDB
	RelationDB
}

// BotDatabase is the database of the bot
// which contains users and other entities
// such as administrator, shoguns with their slaves and etc.
type BotDatabase struct {
	db *pgxpool.Pool
}

// New creates new tgbot database object
// and sets the postgres database connection pool
// to work correctly with concurrent r/w
func New(ctx context.Context, dbConfig *config.Config) *BotDatabase {
	connectionPattern := "postgresql://%s:%s@%s:%s/%s"
	connURL := fmt.Sprintf(connectionPattern,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	connection, _ := pgxpool.New(ctx, connURL)
	return &BotDatabase{db: connection}
}
