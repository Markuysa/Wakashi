package service

import (
	"context"
	"tgBotIntern/app/internal/database"
)

type RelationsServiceMethods interface {
	Bind(ctx context.Context, masterUsername, slaveUsername string) error
}

type RelationsService struct {
	relationDB database.RelationDB
}

func (s *RelationsService) Bind(ctx context.Context, masterUsername, slaveUsername string) error {
	return s.relationDB.BindEntities(ctx, masterUsername, slaveUsername)
}
