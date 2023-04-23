package service

import (
	"context"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
)

type CardRights interface {
	BindToDaimyo(ctx context.Context, cardNumber, ownerID int) error
	CreateCard(ctx context.Context, card entity.Card) error
	GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error)
	SetTotal(ctx context.Context, total, cardNumber int) error
}

type CardService struct {
	cardsRepos database.CardsDB
}

func (s *CardService) CreateCard(ctx context.Context, card entity.Card) error {
	return s.cardsRepos.AddCard(ctx, card)
}
func (s *CardService) BindToDaimyo(ctx context.Context, cardNumber, ownerID int) error {
	return s.cardsRepos.BindCard(ctx, cardNumber, ownerID)
}
func (s *CardService) GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error) {
	return s.cardsRepos.GetCardsList(ctx, ownerID)
}
func (s *CardService) SetTotal(ctx context.Context, total, cardNumber int) error {
	return s.cardsRepos.SetCardTotal(ctx, total, cardNumber)
}
