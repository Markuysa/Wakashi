package usersService

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/domain"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"time"
)

type UsersRepositoryService interface {
	RegisterUser(ctx context.Context, username, nickname, password string, roleID int) error
	AuthorizeUser(ctx context.Context, username, nickname, password string) (tokenService.Tokens, error)
	GetRoleID(ctx context.Context, username string) (int, error)
	CreateUserSession(ctx context.Context, username string, roleID int) (tokenService.Tokens, error)
	IsUserSessionValid(ctx context.Context, username string, roleID int) (bool, error)
	GetSlavesList(ctx context.Context, masterUsername string, slaveRole int) ([]entity.User, error)
	GetUser(ctx context.Context, username string) (*entity.User, error)
	GetUserID(ctx context.Context, username string) (int, error)
	UpdatePassword(ctx context.Context, name string, password string) error
}

type UsersService struct {
	repos           database.UsersDatabase
	tokenManager    tokenService.TokenManager
	refreshTokenTTL time.Duration
	accessTokenTTL  time.Duration
}

func NewUsersService(repos database.UsersDatabase, tokenManager tokenService.TokenManager, refreshTokenTTL time.Duration, accessTokenTTL time.Duration) *UsersService {
	return &UsersService{repos: repos, tokenManager: tokenManager, refreshTokenTTL: refreshTokenTTL, accessTokenTTL: accessTokenTTL}
}
func (u *UsersService) UpdatePassword(ctx context.Context, name string, password string) error {
	return u.repos.UpdatePassword(ctx, name, password)
}
func (u *UsersService) GetUserID(ctx context.Context, username string) (int, error) {
	return u.repos.GetUserID(ctx, username)
}
func (u *UsersService) GetUser(ctx context.Context, username string) (*entity.User, error) {
	return u.repos.GetUser(ctx, username)
}
func (u *UsersService) GetSlavesList(ctx context.Context, masterUsername string, slaveRole int) ([]entity.User, error) {
	userID, err := u.GetUserID(ctx, masterUsername)
	if err != nil {
		return nil, err
	}
	return u.repos.GetSlavesList(ctx, userID, slaveRole)
}

func (u *UsersService) IsUserSessionValid(ctx context.Context, username string, role int) (bool, error) {
	session, err := u.tokenManager.GetUserSession(ctx, username)
	if session != nil {
		return u.tokenManager.ParseToken(ctx, session.AccessToken, username, role)
	}
	return false, errors.New("error with session: token not found", err)
}

func (u *UsersService) CreateUserSession(ctx context.Context, username string, roleID int) (tokenService.Tokens, error) {
	var (
		res tokenService.Tokens
		err error
	)
	res.AccessToken, err = u.tokenManager.NewJWT(username, roleID, u.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = u.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.Session{
		AccessToken: res.AccessToken,
		ExpiresAt:   time.Now().Add(u.refreshTokenTTL),
	}
	err = u.tokenManager.SetUserSession(ctx, username, session)
	if err != nil {
		return tokenService.Tokens{}, err
	}
	if err != nil {
		return tokenService.Tokens{}, err
	}
	return res, nil
}

func (u *UsersService) AuthorizeUser(ctx context.Context, username, nickname, password string) (tokenService.Tokens, error) {
	user, err := u.repos.IsExist(ctx, username, nickname, password)
	if err != nil {
		return tokenService.Tokens{}, err
	}
	return u.CreateUserSession(ctx, username, user.Role)
}

func (u *UsersService) GetRoleID(ctx context.Context, username string) (int, error) {
	return u.repos.GetUserRoleID(ctx, username)
}
func (u *UsersService) RegisterUser(ctx context.Context, username, nickname, password string, roleID int) error {
	user, err := u.repos.GetUser(ctx, username)
	if user != nil {
		return errors.New("user already exists")
	}
	if roleID == -1 {
		return errors.New("error with role: not found")
	}
	err = u.repos.AddUser(ctx, password, username, nickname, roleID)
	if err != nil {
		return errors.New("failed to register user:%v", err)
	}
	return nil
}
