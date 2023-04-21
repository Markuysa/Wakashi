package tokenDb

import (
	"context"
	"encoding/json"
	"errors"
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
	Save(ctx context.Context, username string, session domain.Session) error
	Get(ctx context.Context, username string) (domain.Session, error)
}

type TokenRepository struct {
	db         *redis.Client
	sessionTTL time.Duration
}

func NewTokenRepository(db *redis.Client, sessionTTL time.Duration) *TokenRepository {
	return &TokenRepository{db: db, sessionTTL: sessionTTL}
}
func (t *TokenRepository) Save(ctx context.Context, username string, session domain.Session) error {
	jsonObj, err := json.Marshal(session)
	if err != nil {
		errors.New("failed to serialize session data")
	}
	return t.db.Set(ctx, username, jsonObj, t.sessionTTL).Err()
}
func (t *TokenRepository) Get(ctx context.Context, username string) (domain.Session, error) {
	jsonStr, err := t.db.Get(ctx, username).Result()
	if err != nil {
		return domain.Session{}, errors.New("failed to get token")
	}
	var session domain.Session
	err = json.Unmarshal([]byte(jsonStr), &session)
	if err != nil {
		return domain.Session{}, errors.New("failed to deserialize session data")
	}
	return session, nil
}
