package service

import (
	"context"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type DaimyoRights interface {
	GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error)
	CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64) error
	GetSamuraiList(ctx context.Context, masterUsername string) ([]entity.User, error)
	SetCardsBalances(ctx context.Context, totalValue float64, cardnumber int) error
	BindShogun(ctx context.Context, shogunUsername, daimyoUsername string) error
}

type DaimyoService struct {
	cardsService     CardRights
	userService      usersService.UsersRepositoryService
	relationsService RelationsServiceMethods
}

func NewDaimyoService(cardsService CardRights, userService usersService.UsersRepositoryService, relationsService RelationsServiceMethods) *DaimyoService {
	return &DaimyoService{cardsService: cardsService, userService: userService, relationsService: relationsService}
}

func (s *DaimyoService) GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error) {
	return s.cardsService.GetCardsList(ctx, ownerID)
}

func (s *DaimyoService) CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64) error {
	return nil
}

func (s *DaimyoService) GetSamuraiList(ctx context.Context, masterUsername string) ([]entity.User, error) {
	return s.userService.GetSlavesList(ctx, masterUsername, roles.Samurai)
}

func (s *DaimyoService) SetCardsBalances(ctx context.Context, totalValue float64, cardnumber int) error {
	return s.cardsService.SetTotal(ctx, totalValue, cardnumber)
}

func (s *DaimyoService) BindShogun(ctx context.Context, shogunUsername, daimyoUsername string) error {
	return s.relationsService.Bind(ctx, shogunUsername, daimyoUsername)
}
