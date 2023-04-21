package tokenService

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"math/rand"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/domain"
	"tgBotIntern/app/pkg/auth/tokenDb"
	"time"
)

type TokenManager interface {
	NewRefreshToken() (string, error)
	NewJWT(username string, role string, ttl time.Duration) (string, error)
	SetUserSession(ctx context.Context, username string, session domain.Session) error
	GetUserSession(ctx context.Context, username string) (*domain.Session, error)
	ParseToken(ctx context.Context, tokenString string) (entity.User, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenService struct {
	TokenRepository tokenDb.TokenRepos
}

func (s *TokenService) SetUserSession(ctx context.Context, username string, session domain.Session) error {
	return s.TokenRepository.Save(ctx, username, session)
}
func NewTokenService(tokenRepository tokenDb.TokenRepos) *TokenService {

	return &TokenService{TokenRepository: tokenRepository}
}
func (s *TokenService) NewJWT(username string, role string, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = role
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

func (s *TokenService) ParseToken(ctx context.Context, tokenString string) (entity.User, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-secret-key"), nil
	})
	if err != nil {
		return entity.User{}, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// print the user info
		user := entity.User{
			Role:     claims["role"].(string),
			Username: claims["username"].(string),
		}
		return user, nil
	} else {
		return entity.User{}, errors.New("invalid token")
	}
}
