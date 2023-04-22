package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type AuthHandler interface {
	HandleRegister(c *gin.Context)
	HandleLogin(c *gin.Context)
	HandleGetLoginPage(c *gin.Context)
	HandleGetRegisterPage(c *gin.Context)
}
type Claims struct {
	jwt.StandardClaims
	Username string `json:"user_id"`
	Role     string `json:"role"`
}

type TgBotAuth struct {
	usersService      *usersService.UsersService
	tokenReposService *tokenService.TokenService
}

func NewTgBotAuth(usersService *usersService.UsersService) *TgBotAuth {
	return &TgBotAuth{usersService: usersService}
}

func (t *TgBotAuth) HandleRegister(c *gin.Context) {

	var user entity.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err := t.usersService.RegisterUser(c, user.Username, user.Password, user.Role)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

}
func (t *TgBotAuth) HandleLogin(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := t.usersService.AuthorizeUser(c, user.Username, user.Password)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (t *TgBotAuth) HandleGetLoginPage(c *gin.Context) {
	html, err := os.ReadFile("app/pkg/auth/ui/html/login.html")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = c.Writer.Write(html)
	if err != nil {
		return
	}
}
func (t *TgBotAuth) HandleGetRegisterPage(c *gin.Context) {
	html, err := os.ReadFile("app/pkg/auth/ui/html/register.html")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = c.Writer.Write(html)
	if err != nil {
		return
	}
}
