package usersService

import (
	"context"
	"reflect"
	"testing"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/tokenService"
	"time"
)

func TestNewUsersService(t *testing.T) {
	type args struct {
		repos        *database.BotDatabase
		tokenManager tokenService.TokenManager
	}
	tests := []struct {
		name string
		args args
		want *UsersService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.repos, tt.args.tokenManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_AuthorizeUser(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    tokenService.Tokens
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			got, err := u.AuthorizeUser(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizeUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_CreateUserSession(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx      context.Context
		username string
		roleID   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    tokenService.Tokens
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			got, err := u.CreateUserSession(tt.args.ctx, tt.args.username, tt.args.roleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUserSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_GetRoleID(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			got, err := u.GetRoleID(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRoleID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_GetSlavesList(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx            context.Context
		masterUsername string
		slaveRole      int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			got, err := u.GetSlavesList(tt.args.ctx, tt.args.masterUsername, tt.args.slaveRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSlavesList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSlavesList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_IsUserSessionValid(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx      context.Context
		username string
		role     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			got, err := u.IsUserSessionValid(tt.args.ctx, tt.args.username, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsUserSessionValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsUserSessionValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersService_RegisterUser(t *testing.T) {
	type fields struct {
		repos           database.TelegramDB
		tokenManager    tokenService.TokenManager
		refreshTokenTTL time.Duration
		accessTokenTTL  time.Duration
	}
	type args struct {
		ctx      context.Context
		username string
		password string
		roleID   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersService{
				repos:           tt.fields.repos,
				tokenManager:    tt.fields.tokenManager,
				refreshTokenTTL: tt.fields.refreshTokenTTL,
				accessTokenTTL:  tt.fields.accessTokenTTL,
			}
			if err := u.RegisterUser(tt.args.ctx, tt.args.username, tt.args.password, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
