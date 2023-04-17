package usersService

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/database"
)

type UsersService struct {
	repos *database.BotDatabase
}

func New(repos *database.BotDatabase) *UsersService {
	return &UsersService{repos: repos}
}

func (u *UsersService) AuthorizeUser(ctx context.Context, username, password string, role int) error {
	exist, err := u.repos.IsExist(ctx, username, password)
	if err != nil {
		return errors.New("failed to authorize user:%v", err)
	}
	if exist {
		role, err := u.repos.GetUserRoleID(ctx, username)
		if err != nil {
			return errors.New("failed to define the role of user:%v", err)
		}
		ctx = context.WithValue(ctx, "ROLE", role)
	} else {
		return errors.New("invalid username or password!")
	}
	return nil
}
