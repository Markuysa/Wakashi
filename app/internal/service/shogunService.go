package service

import (
	"context"
	"errors"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

// ShogunRights is an interface that represents the methods of
// ShogunService structure - the use case layer of program to interact with shogun entities
// The GetSlavesList method is used to get slaves list of shogun (daimyo, samurai)
// The BindCardToDaimyo method is used to bind card to daimyo
type ShogunRights interface {
	GetSlavesList(ctx context.Context, masterUsername string, slaveRole int) ([]entity.User, error)
	CreateCard(ctx context.Context, card entity.Card) error
	BindCardToDaimyo(ctx context.Context, cardNumber string, daimyoUsername string) error
}

type ShogunService struct {
	usersService usersService.UsersRepositoryService
	cardService  CardRights
}

func NewShogunService(usersService usersService.UsersRepositoryService, cardService CardRights) *ShogunService {
	return &ShogunService{usersService: usersService, cardService: cardService}
}

// Add some validation
func (s *ShogunService) GetSlavesList(ctx context.Context, masterUsername string, slaveRole int) ([]entity.User, error) {
	return s.usersService.GetSlavesList(ctx, masterUsername, slaveRole)
}
func (s *ShogunService) CreateCard(ctx context.Context, card entity.Card) error {
	// ignore the error because it appears when the card with that card_number doesnt exist
	cardExists, _ := s.cardService.GetCard(ctx, int(card.CardNumber))
	if cardExists != nil {
		return errors.New("card with that card_number already exists")
	}
	return s.cardService.CreateCard(ctx, card)
}
func (s *ShogunService) BindCardToDaimyo(ctx context.Context, cardNumber string, daimyoUsername string) error {
	return s.cardService.BindToDaimyo(ctx, cardNumber, daimyoUsername)
}
