package service

import (
	"context"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

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

func (s *ShogunService) GetSlavesList(ctx context.Context, masterUsername string, slaveRole int) ([]entity.User, error) {
	return s.usersService.GetSlavesList(ctx, masterUsername, slaveRole)
}
func (s *ShogunService) CreateCard(ctx context.Context, card entity.Card) error {
	return s.cardService.CreateCard(ctx, card)
}
func (s *ShogunService) BindCardToDaimyo(ctx context.Context, cardNumber string, daimyoUsername string) error {
	return s.cardService.BindToDaimyo(ctx, cardNumber, daimyoUsername)
}
