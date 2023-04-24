package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"tgBotIntern/app/internal/database/config"
)

// NewDBConnection creates new tgbot database object
// and sets the postgres database connection pool
// to work correctly with concurrent r/w
func NewDBConnection(ctx context.Context, dbConfig *config.Config) *pgxpool.Pool {
	connectionPattern := "postgresql://%s:%s@%s:%s/%s"
	connURL := fmt.Sprintf(connectionPattern,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	connection, _ := pgxpool.New(ctx, connURL)
	return connection
}
