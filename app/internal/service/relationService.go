package service

import (
	"context"
	"errors"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

// RelationsServiceMethods is an interface that represents the methods of
// RelationsService structure - the use case layer of program to interact with relations
// btwn entities
// The Bind method is used to bind two entities with one another
// The GetMasterUsername method is used to get master's username by slave's username
type RelationsServiceMethods interface {
	Bind(ctx context.Context, masterUsername, slaveUsername string) error
	GetMasterUsername(ctx context.Context, slaveUsername string) (string, error)
}

type RelationsService struct {
	relationDB   database.RelationDatabase
	usersService usersService.UsersRepositoryService
}

func NewRelationsService(relationDB database.RelationDatabase, usersService usersService.UsersRepositoryService) *RelationsService {
	return &RelationsService{relationDB: relationDB, usersService: usersService}
}
func (s *RelationsService) GetMasterUsername(ctx context.Context, slaveUsername string) (string, error) {
	slaveID, err := s.usersService.GetUserID(ctx, slaveUsername)
	if err != nil {
		return "", err
	}
	return s.relationDB.GetMasterUsername(ctx, slaveID)

}
func (s *RelationsService) Bind(ctx context.Context, masterUsername, slaveUsername string) error {
	masterID, err := s.usersService.GetUserID(ctx, masterUsername)
	if err != nil {
		return errors.New("No user with that username: " + masterUsername)
	}
	slaveID, err := s.usersService.GetUserID(ctx, slaveUsername)
	if err != nil {
		return errors.New("No user with that username: " + slaveUsername)
	}

	return s.relationDB.BindEntities(ctx, masterID, slaveID)
}
