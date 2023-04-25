package service

import (
	"context"
	"errors"
	"fmt"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
	"time"
)

type DaimyoRights interface {
	GetCardsList(ctx context.Context, ownerUsername string) ([]entity.Card, error)
	CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64, ownerUsername string) error
	GetSamuraiList(ctx context.Context, masterUsername string) ([]entity.User, error)
	SetCardsBalances(ctx context.Context, totalValue float64, cardnumber int) error
	BindShogun(ctx context.Context, shogunUsername, daimyoUsername string) error
	GetSamuraiTurnover(ctx context.Context, masterUsername, username string) (float64, error)
	GetCardsTotal(ctx context.Context, username string) (float64, error)
}

type DaimyoService struct {
	cardsService        CardRights
	userService         usersService.UsersRepositoryService
	relationsService    RelationsServiceMethods
	transactionsService TransactionProcessor
}

func NewDaimyoService(cardsService CardRights, userService usersService.UsersRepositoryService, relationsService RelationsServiceMethods, transactionsService TransactionProcessor) *DaimyoService {
	return &DaimyoService{cardsService: cardsService, userService: userService, relationsService: relationsService, transactionsService: transactionsService}
}
func (s *DaimyoService) GetSamuraiTurnover(ctx context.Context, masterUsername, username string) (float64, error) {
	turnover, err := s.cardsService.GetTurnover(ctx, masterUsername, username)
	if err != nil {
		return 0, errors.New("cant get the turnover of user: " + username)
	}
	return turnover, err
}
func (s *DaimyoService) GetCardsTotal(ctx context.Context, username string) (float64, error) {
	return s.cardsService.GetTotal(ctx, username)
}
func (s *DaimyoService) GetCardsList(ctx context.Context, ownerUsername string) ([]entity.Card, error) {
	userID, err := s.userService.GetUserID(ctx, ownerUsername)
	if err != nil {
		return nil, errors.New("failed to find user: " + ownerUsername)
	}
	return s.cardsService.GetCardsList(ctx, userID)
}

func (s *DaimyoService) CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64, ownerUsername string) error {
	userID, err := s.userService.GetUserID(ctx, ownerUsername)
	if err != nil {
		return errors.New("User: " + ownerUsername + " not found")
	}
	card, err := s.cardsService.GetCard(ctx, cardID)
	if card == nil || err != nil {
		return errors.New(fmt.Sprintf("Erros with getting informations about your card with id:%v ", cardID))
	}
	return s.transactionsService.CreateTransactionRequest(ctx, entity.Transaction{
		CardNumber:      cardID,
		OwnerID:         userID,
		OperationValue:  increasementValue,
		TransactionDate: time.Now(),
		Status:          false,
	})
}

func (s *DaimyoService) GetSamuraiList(ctx context.Context, masterUsername string) ([]entity.User, error) {
	return s.userService.GetSlavesList(ctx, masterUsername, roles.Samurai)
}

func (s *DaimyoService) SetCardsBalances(ctx context.Context, totalValue float64, cardnumber int) error {
	return s.cardsService.SetTotal(ctx, totalValue, cardnumber)
}

func (s *DaimyoService) BindShogun(ctx context.Context, shogunUsername, daimyoUsername string) error {
	roleID, err := s.userService.GetRoleID(ctx, shogunUsername)
	if err != nil {
		return errors.New("cannot find user with username: " + shogunUsername)
	}
	if roleID != roles.Shogun {
		return errors.New("The user: " + shogunUsername + " is not a shogun!")
	}
	return s.relationsService.Bind(ctx, shogunUsername, daimyoUsername)
}
