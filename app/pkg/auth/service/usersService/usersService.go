package usersService

import (
	"context"
	"fmt"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/constants/roles"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/pkg/auth/domain"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"time"
)

type UsersRepositoryService interface {
	RegisterUser(ctx context.Context, username, password string, role string) error
	AuthorizeUser(ctx context.Context, username, password string) (tokenService.Tokens, error)
	GetRole(ctx context.Context, username string) (int, error)
	CreateUserSession(ctx context.Context, username string) (tokenService.Tokens, error)
}

type UsersService struct {
	repos           *database.BotDatabase
	tokenManager    tokenService.TokenManager
	refreshTokenTTL time.Duration
	accessTokenTTL  time.Duration
}

func NewUsersService(repos *database.BotDatabase, tokenManager tokenService.TokenManager) *UsersService {
	return &UsersService{repos: repos, tokenManager: tokenManager}
}
func (u *UsersService) CreateUserSession(ctx context.Context, username string) (tokenService.Tokens, error) {
	var (
		res tokenService.Tokens
		err error
	)
	res.AccessToken, err = u.tokenManager.NewJWT(username, u.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = u.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(u.refreshTokenTTL),
	}

	err = u.repos.SetUserSession(ctx, username, session)
	return res, err
}

func (u *UsersService) AuthorizeUser(ctx context.Context, username, password string) (tokenService.Tokens, error) {
	exist, err := u.repos.IsExist(ctx, username, password)
	if err != nil {
		return tokenService.Tokens{}, errors.New("failed to authorize user:%v", err)
	}
	if exist {
		role, err := u.repos.GetUserRoleID(ctx, username)
		if err != nil {
			return tokenService.Tokens{}, errors.New("failed to define the role of user:%v", err)
		}
		fmt.Println(role)
		// check if the token is correct
	} else {
		return tokenService.Tokens{}, errors.New("invalid username or password!")
	}
	return u.CreateUserSession(ctx, username)
}

func (u *UsersService) GetRole(ctx context.Context, username string) (int, error) {
	return u.repos.GetUserRoleID(ctx, username)
}
func (u *UsersService) RegisterUser(ctx context.Context, username, password string, role string) error {
	exist, err := u.repos.IsExist(ctx, username, password)
	if err != nil {
		return errors.New("failed to register user:%v", err)
	}
	if !exist {
		return errors.New("user already exists")
	}
	err = u.repos.AddUser(ctx, password, username, roles.GetRoleString(role))
	if err != nil {
		return errors.New("failed to register user:%v", err)
	}
	return nil
}
