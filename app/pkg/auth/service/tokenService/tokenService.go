package tokenService

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"math/rand"
	"tgBotIntern/app/pkg/auth/tokenDb"
	"time"
)

type TokenManager interface {
	NewRefreshToken() (string, error)
	NewJWT(username string, ttl time.Duration) (string, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type TokenService struct {
	TokenRepository tokenDb.TokenRepos
}

func NewTokenService(tokenRepository tokenDb.TokenRepos) *TokenService {

	return &TokenService{TokenRepository: tokenRepository}
}
func (s *TokenService) NewJWT(username string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(ttl)),
		Subject:   username,
	})
	var somestring string
	return token.SignedString([]byte(somestring))
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
