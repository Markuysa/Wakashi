package tokenService

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"math/rand"
	"tgBotIntern/app/pkg/auth/database"
	"tgBotIntern/app/pkg/auth/domain"
	"time"
)

type TokenManager interface {
	NewRefreshToken() (string, error)
	NewJWT(username string, role int, ttl time.Duration) (string, error)
	SetUserSession(ctx context.Context, username string, session domain.Session) error
	GetUserSession(ctx context.Context, username string) (*domain.Session, error)
	ParseToken(ctx context.Context, tokenString, username string, role int) (bool, error)
	SetCurrentActiveUser(ctx context.Context, username string) error
	ResetUserSession(ctx context.Context, username string) error
	IsTokenValid(ctx context.Context, token, username string) (bool, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenService struct {
	TokenRepository database.TokenRepos
}

func (s *TokenService) IsTokenValid(ctx context.Context, token, username string) (bool, error) {

	tokenEntry, err := s.TokenRepository.Get(ctx, username)
	if err != nil {
		return false, errors.New("No user with that username: " + username)
	}
	if tokenEntry.AccessToken == token {
		return true, nil
	}
	return false, errors.New("invalid token: " + token)

}

func (s *TokenService) ResetUserSession(ctx context.Context, username string) error {
	return s.TokenRepository.Remove(ctx, username)
}

func (s *TokenService) SetCurrentActiveUser(ctx context.Context, username string) error {
	return s.TokenRepository.SaveCurrentUser(ctx, username)
}

func (s *TokenService) SetUserSession(ctx context.Context, username string, session domain.Session) error {
	return s.TokenRepository.SaveSession(ctx, username, session)
}
func NewTokenService(tokenRepository database.TokenRepos) *TokenService {

	return &TokenService{TokenRepository: tokenRepository}
}
func (s *TokenService) NewJWT(username string, roleID int, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = roleID
	claims["username"] = username
	claims["expire"] = ttl
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (s *TokenService) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
func (s *TokenService) GetUserSession(ctx context.Context, username string) (*domain.Session, error) {
	get, err := s.TokenRepository.Get(ctx, username)
	if err != nil {
		return nil, err
	}
	return &get, nil
}

func (s *TokenService) ParseToken(ctx context.Context, tokenString, username string, role int) (bool, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-secret-key"), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		roleClaims := int(claims["role"].(float64))
		usernameClaims := claims["username"].(string)
		if usernameClaims == username && roleClaims == role {
			return true, nil
		}
		return false, nil
	} else {
		return false, errors.New("invalid token")
	}
}
