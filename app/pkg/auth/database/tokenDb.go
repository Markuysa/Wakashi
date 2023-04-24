package database

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"tgBotIntern/app/pkg/auth/config"
	"tgBotIntern/app/pkg/auth/domain"
	"time"
)

type TokenRepos interface {
	SaveSession(ctx context.Context, username string, session domain.Session) error
	SaveCurrentUser(ctx context.Context, username string) error
	Get(ctx context.Context, username string) (domain.Session, error)
	Remove(ctx context.Context, key string) error
}

type TokenRepository struct {
	db         *redis.Client
	sessionTTL time.Duration
}

func NewTokenRepository(config config.Config) *TokenRepository {
	sessionTTL := 12 * time.Hour
	client := redis.NewClient(&redis.Options{
		Addr: config.Host + ":" + config.Port,
		DB:   0,
	})
	return &TokenRepository{db: client, sessionTTL: sessionTTL}
}

func (t *TokenRepository) Remove(ctx context.Context, key string) error {
	return t.db.Del(ctx, key).Err()
}

func (t *TokenRepository) SaveCurrentUser(ctx context.Context, username string) error {
	return t.db.Set(ctx, "activeUser", username, time.Hour).Err()
}
func (t *TokenRepository) SaveSession(ctx context.Context, username string, session domain.Session) error {
	jsonObj, err := json.Marshal(session)
	if err != nil {
		return errors.New("failed to serialize session data")
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
