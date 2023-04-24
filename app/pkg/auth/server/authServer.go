package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tgBotIntern/app/pkg/auth/database"
	"tgBotIntern/app/pkg/auth/handlers"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

// its better to move server to independent service

type Server interface {
	setHandlers()
	Start() error
}

type AuthSerer struct {
	port            string
	ServerEngine    *gin.Engine
	RequestHandlers handlers.AuthHandler
}

func NewAuthSerer(port string, repository *database.TokenRepository, service *usersService.UsersService) *AuthSerer {
	return &AuthSerer{
		port:            port,
		ServerEngine:    gin.New(),
		RequestHandlers: handlers.NewTgBotAuth(service),
	}
}
func (a *AuthSerer) setHandlers() {
	a.ServerEngine.Handle(http.MethodPost, "/register", a.RequestHandlers.HandleRegister)
	a.ServerEngine.Handle(http.MethodPost, "/login", a.RequestHandlers.HandleLogin)
	a.ServerEngine.Handle(http.MethodGet, "/register", a.RequestHandlers.HandleGetRegisterPage)
	a.ServerEngine.Handle(http.MethodGet, "/login", a.RequestHandlers.HandleGetLoginPage)
}
func (a *AuthSerer) Start() error {
	a.setHandlers()
	return a.ServerEngine.Run(a.port)
}
