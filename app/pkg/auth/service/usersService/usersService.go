package usersService

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/pkg/auth/domain"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"time"
)

type UsersRepositoryService interface {
	RegisterUser(ctx context.Context, username, password string, roleID int) error
	AuthorizeUser(ctx context.Context, username, password string) (tokenService.Tokens, error)
	GetRoleID(ctx context.Context, username string) (int, error)
	CreateUserSession(ctx context.Context, username string, roleID int) (tokenService.Tokens, error)
	IsUserSessionValid(ctx context.Context, username string, roleID int) (bool, error)
}

type UsersService struct {
	repos           database.TelegramDB
	tokenManager    tokenService.TokenManager
	refreshTokenTTL time.Duration
	accessTokenTTL  time.Duration
}

func NewUsersService(repos *database.BotDatabase, tokenManager tokenService.TokenManager) *UsersService {
	return &UsersService{repos: repos, tokenManager: tokenManager}
}
func (u *UsersService) IsUserSessionValid(ctx context.Context, username string, role int) (bool, error) {
	session, err := u.tokenManager.GetUserSession(ctx, username)
	if session != nil {
		user, err := u.tokenManager.ParseToken(ctx, session.AccessToken)
		if err != nil {
			return false, err
		}
		if user.Role == role {
			return true, nil
		}
		return false, err
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
	return res, err
}

func (u *UsersService) AuthorizeUser(ctx context.Context, username, password string) (tokenService.Tokens, error) {
	user, err := u.repos.IsExist(ctx, username, password)
	if err != nil {
		return tokenService.Tokens{}, errors.New("failed to authorize user:%v", err)
	}
	return u.CreateUserSession(ctx, username, user.Role)
}

func (u *UsersService) GetRoleID(ctx context.Context, username string) (int, error) {
	return u.repos.GetUserRoleID(ctx, username)
}
func (u *UsersService) RegisterUser(ctx context.Context, username, password string, roleID int) error {
	user, err := u.repos.GetUser(ctx, username)
	if user != nil {
		return errors.New("user already exists")
	}
	if roleID == -1 {
		return errors.New("error with role: not found")
	}
	err = u.repos.AddUser(ctx, password, username, roleID)
	if err != nil {
		return errors.New("failed to register user:%v", err)
	}
	return nil
}
