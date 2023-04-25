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

// SamuraiRights is an interface that represents the methods of
// SamuraiService structure - the use case layer of program to interact with samurai entities
// The GetTurnover method is used to get turnover of the samurai
// The BindToDaimyo method is used to bind samurai to daimyo
// The CreateCardIncreasementRequest method is used to create request to collector to update
// card balance
type SamuraiRights interface {
	GetTurnover(ctx context.Context, samuraiUsername string) (float64, error)
	BindToDaimyo(ctx context.Context, masterUsername, slaveUsername string) error
	CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64, ownerUsername string) error
}

type SamuraiService struct {
	relationService     RelationsServiceMethods
	cardsService        CardRights
	usersService        usersService.UsersRepositoryService
	transactionsService TransactionProcessor
}

func (s *SamuraiService) CreateCardIncreasementRequest(ctx context.Context, cardID int, increasementValue float64, ownerUsername string) error {
	userID, err := s.usersService.GetUserID(ctx, ownerUsername)
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

func NewSamuraiService(relationService RelationsServiceMethods, cardsService CardRights, usersService usersService.UsersRepositoryService) *SamuraiService {
	return &SamuraiService{relationService: relationService, cardsService: cardsService, usersService: usersService}
}

func (s *SamuraiService) GetTurnover(ctx context.Context, samuraiUsername string) (float64, error) {
	userRole, err := s.usersService.GetRoleID(ctx, samuraiUsername)
	if err != nil {
		return 0, errors.New("No user with username: " + samuraiUsername)
	}
	if userRole != roles.Samurai {
		return 0, errors.New("The user: " + samuraiUsername + " is not a Samurai!")
	}
	masterUsername, err := s.relationService.GetMasterUsername(ctx, samuraiUsername)
	if err != nil {
		return 0, errors.New("failed to get master of " + samuraiUsername)
	}
	return s.cardsService.GetTurnover(ctx, masterUsername, samuraiUsername)
}

func (s *SamuraiService) BindToDaimyo(ctx context.Context, masterUsername, slaveUsername string) error {
	return s.relationService.Bind(ctx, masterUsername, slaveUsername)
}
