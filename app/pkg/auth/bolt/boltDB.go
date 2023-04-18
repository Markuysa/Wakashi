package bolt

import (
	"errors"
	"go.etcd.io/bbolt"
	"strconv"
	"tgBotIntern/app/pkg/auth"
)

type TokenRepository struct {
	db *bbolt.DB
}

func NewTokenRepository(db *bbolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (t *TokenRepository) Save(chatID int64, token string, bucket auth.Bucket) error {
	return t.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(strconv.FormatInt(chatID, 10)), []byte(token))
	})
}
func (t *TokenRepository) Get(chatID int64, bucket auth.Bucket) (string, error) {
	var token string
	err := t.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get([]byte(strconv.FormatInt(chatID, 10)))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("can't find the token")
	}
	return token, nil
}
