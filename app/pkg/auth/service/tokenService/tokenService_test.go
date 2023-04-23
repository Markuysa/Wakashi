package tokenService

import (
	"context"
	"reflect"
	"testing"
	"tgBotIntern/app/pkg/auth/domain"
	"tgBotIntern/app/pkg/auth/tokenDb"
	"time"
)

func TestNewTokenService(t *testing.T) {
	type args struct {
		tokenRepository tokenDb.TokenRepos
	}
	tests := []struct {
		name string
		args args
		want *TokenService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenService(tt.args.tokenRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_GetUserSession(t *testing.T) {
	type fields struct {
		TokenRepository tokenDb.TokenRepos
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenService{
				TokenRepository: tt.fields.TokenRepository,
			}
			got, err := s.GetUserSession(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_NewJWT(t *testing.T) {
	type fields struct {
		TokenRepository tokenDb.TokenRepos
	}
	type args struct {
		username string
		roleID   int
		ttl      time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenService{
				TokenRepository: tt.fields.TokenRepository,
			}
			got, err := s.NewJWT(tt.args.username, tt.args.roleID, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_NewRefreshToken(t *testing.T) {
	type fields struct {
		TokenRepository tokenDb.TokenRepos
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenService{
				TokenRepository: tt.fields.TokenRepository,
			}
			got, err := s.NewRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewRefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_ParseToken(t *testing.T) {
	type fields struct {
		TokenRepository tokenDb.TokenRepos
	}
	type args struct {
		ctx         context.Context
		tokenString string
		username    string
		role        int
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
			s := &TokenService{
				TokenRepository: tt.fields.TokenRepository,
			}
			got, err := s.ParseToken(tt.args.ctx, tt.args.tokenString, tt.args.username, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenService_SetUserSession(t *testing.T) {
	type fields struct {
		TokenRepository tokenDb.TokenRepos
	}
	type args struct {
		ctx      context.Context
		username string
		session  domain.Session
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
			s := &TokenService{
				TokenRepository: tt.fields.TokenRepository,
			}
			if err := s.SetUserSession(tt.args.ctx, tt.args.username, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("SetUserSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
