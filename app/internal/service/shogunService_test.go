package service

import (
	"context"
	"reflect"
	"testing"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

func TestShogunService_BindCardToDaimyo(t *testing.T) {
	type fields struct {
		usersService usersService.UsersService
		cardService  CardService
	}
	type args struct {
		ctx        context.Context
		cardNumber int
		daimyoID   int
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
			s := &ShogunService{
				usersService: tt.fields.usersService,
				cardService:  tt.fields.cardService,
			}
			if err := s.BindCardToDaimyo(tt.args.ctx, tt.args.cardNumber, tt.args.daimyoID); (err != nil) != tt.wantErr {
				t.Errorf("BindCardToDaimyo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShogunService_CreateCard(t *testing.T) {
	type fields struct {
		usersService usersService.UsersService
		cardService  CardService
	}
	type args struct {
		ctx  context.Context
		card entity.Card
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
			s := &ShogunService{
				usersService: tt.fields.usersService,
				cardService:  tt.fields.cardService,
			}
			if err := s.CreateCard(tt.args.ctx, tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShogunService_GetSlavesList(t *testing.T) {
	type fields struct {
		usersService usersService.UsersService
		cardService  CardService
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
			s := &ShogunService{
				usersService: tt.fields.usersService,
				cardService:  tt.fields.cardService,
			}
			got, err := s.GetSlavesList(tt.args.ctx, tt.args.masterUsername, tt.args.slaveRole)
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
