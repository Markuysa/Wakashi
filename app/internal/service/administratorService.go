package service

import (
	"context"
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
	BindCardToDaimyo(ctx context.Context, daimyoID, cardNumber int) error
	GetEntityReport(ctx context.Context, entityID int) error
}

type AdministratorService struct {
	usersService    usersService.UsersService
	cardService     CardService
	relationService RelationsServiceMethods
}

func (a *AdministratorService) CreateEntity(ctx context.Context, user entity.User) error {
	return a.usersService.RegisterUser(ctx, user.Username, user.Password, user.Role)
}

func (a *AdministratorService) CreateCard(ctx context.Context, card entity.Card) error {
	return a.cardService.CreateCard(ctx, card)
}
func (a *AdministratorService) BindCardToDaimyo(ctx context.Context, daimyoID, cardNumber int) error {
	return a.cardService.BindToDaimyo(ctx, cardNumber, daimyoID)
}

func (a *AdministratorService) BindSlave(ctx context.Context, masterUsername, slaveUsername string) error {
	return a.relationService.Bind(ctx, masterUsername, slaveUsername)
}
func (a *AdministratorService) GetEntityReport(ctx context.Context, entityID int) error {
	return nil
}
