package domain

import "time"

type Session struct {
	AccessToken string    `json:"accessToken" bson:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt" bson:"expiresAt"`
}
