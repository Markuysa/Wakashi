package tokenDb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"tgBotIntern/app/pkg/auth/domain"
	"time"
)

type Bucket string

const (
	AccessTokens Bucket = "access_token"
)

// change to redis
type TokenRepos interface {
	Save(ctx context.Context, username string, session domain.Session)
	Get(ctx context.Context, username string) (domain.Session, error)
}

type TokenRepository struct {
	db         *redis.Client
	sessionTTL time.Duration
}

func NewTokenRepository(db *redis.Client, sessionTTL time.Duration) *TokenRepository {
	return &TokenRepository{db: db, sessionTTL: sessionTTL}
}
func (t *TokenRepository) Save(ctx context.Context, username string, session domain.Session) {
	t.db.Set(ctx, username, session, t.sessionTTL)
}
func (t *TokenRepository) Get(ctx context.Context, username string) (domain.Session, error) {
	var session domain.Session
	err := t.db.Get(ctx, username).Scan(&session)
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil
}
