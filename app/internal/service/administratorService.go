package service

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

// AdministratorRights is an interface which represents
// the methods available to the administrator
// CreateEntity method crates new entity object
// CreateCard method created new bank card
// BindSlave method binds slave object with master object
// BindCardToDaimyo method binds a new card to Daimyo
// GetEntityReport method creates full information about any entity
type AdministratorRights interface {
	CreateEntity(ctx context.Context, user entity.User) error
	CreateCard(ctx context.Context, card entity.Card) error
	BindSlave(ctx context.Context, masterUsername, slaveUsername string) error
	BindCardToDaimyo(ctx context.Context, daimyoUsername, cardNumber string) error
	GetEntityReport(ctx context.Context, userName string) (*entity.Report, error)
}

type AdministratorService struct {
	usersService        usersService.UsersRepositoryService
	cardService         CardRights
	relationService     RelationsServiceMethods
	transactionsService TransactionProcessor
}

func NewAdministratorService(usersService usersService.UsersRepositoryService, cardService CardRights, relationService RelationsServiceMethods, transactionsService TransactionProcessor) *AdministratorService {
	return &AdministratorService{usersService: usersService, cardService: cardService, relationService: relationService, transactionsService: transactionsService}
}

func (a *AdministratorService) CreateEntity(ctx context.Context, user entity.User) error {
	return a.usersService.RegisterUser(ctx, user.Username, user.Password, user.Role)
}

func (a *AdministratorService) CreateCard(ctx context.Context, card entity.Card) error {
	return a.cardService.CreateCard(ctx, card)
}
func (a *AdministratorService) BindCardToDaimyo(ctx context.Context, daimyoUsername, cardNumber string) error {
	return a.cardService.BindToDaimyo(ctx, cardNumber, daimyoUsername)
}

func (a *AdministratorService) BindSlave(ctx context.Context, masterUsername, slaveUsername string) error {
	return a.relationService.Bind(ctx, masterUsername, slaveUsername)
}
func (a *AdministratorService) GetEntityReport(ctx context.Context, userName string) (*entity.Report, error) {
	user, err := a.usersService.GetUser(ctx, userName)
	if err != nil {
		return nil, errors.New("failed to get entity report: %v", err)
	}
	switch user.Role {
	// maybe calculate turnover of shogun
	case roles.Administrator:
		{
			return entity.NewReport(
				user.Username, user.Role, []entity.Transaction{},
			), nil
		}
	case roles.Shogun:
		{
			return entity.NewReport(
				user.Username, user.Role, []entity.Transaction{},
			), nil
		}
	case roles.Daimyo:
		{
			transactionsList, err := a.transactionsService.GetTransactions(ctx, userName)
			if err != nil {
				return nil, err
			}
			return entity.NewReport(
				user.Username, user.Role, transactionsList,
			), nil
		}
	case roles.Samurai:
		{
			transactionsList, err := a.transactionsService.GetTransactions(ctx, userName)
			if err != nil {
				return nil, err
			}
			return entity.NewReport(
				user.Username, user.Role, transactionsList,
			), nil
		}
	default:
		return &entity.Report{}, nil
	}

}
