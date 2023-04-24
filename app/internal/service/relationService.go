package service

import (
	"context"
	"errors"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type RelationsServiceMethods interface {
	Bind(ctx context.Context, masterUsername, slaveUsername string) error
}

type RelationsService struct {
	relationDB   database.RelationDatabase
	usersService usersService.UsersRepositoryService
}

func NewRelationsService(relationDB database.RelationDatabase, usersService usersService.UsersRepositoryService) *RelationsService {
	return &RelationsService{relationDB: relationDB, usersService: usersService}
}

func (s *RelationsService) Bind(ctx context.Context, masterUsername, slaveUsername string) error {
	masterID, err := s.usersService.GetRoleID(ctx, masterUsername)
	if err != nil {
		return errors.New("No user with that username: " + masterUsername)
	}
	slaveID, err := s.usersService.GetRoleID(ctx, slaveUsername)
	if err != nil {
		return errors.New("No user with that username: " + slaveUsername)
	}

	return s.relationDB.BindEntities(ctx, masterID, slaveID)
}
