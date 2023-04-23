package service

import (
	"context"
	"testing"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

func TestAdministratorService_BindCardToDaimyo(t *testing.T) {
	type fields struct {
		usersService    usersService.UsersService
		cardService     CardService
		relationService RelationsServiceMethods
	}
	type args struct {
		ctx        context.Context
		daimyoID   int
		cardNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "no daimyo with that name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdministratorService{
				usersService:    tt.fields.usersService,
				cardService:     tt.fields.cardService,
				relationService: tt.fields.relationService,
			}
			if err := a.BindCardToDaimyo(tt.args.ctx, tt.args.daimyoID, tt.args.cardNumber); (err != nil) != tt.wantErr {
				t.Errorf("BindCardToDaimyo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdministratorService_BindSlave(t *testing.T) {
	type fields struct {
		usersService    usersService.UsersService
		cardService     CardService
		relationService RelationsServiceMethods
	}
	type args struct {
		ctx            context.Context
		masterUsername string
		slaveUsername  string
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
			a := &AdministratorService{
				usersService:    tt.fields.usersService,
				cardService:     tt.fields.cardService,
				relationService: tt.fields.relationService,
			}
			if err := a.BindSlave(tt.args.ctx, tt.args.masterUsername, tt.args.slaveUsername); (err != nil) != tt.wantErr {
				t.Errorf("BindSlave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdministratorService_CreateCard(t *testing.T) {
	type fields struct {
		usersService    usersService.UsersService
		cardService     CardService
		relationService RelationsServiceMethods
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
			a := &AdministratorService{
				usersService:    tt.fields.usersService,
				cardService:     tt.fields.cardService,
				relationService: tt.fields.relationService,
			}
			if err := a.CreateCard(tt.args.ctx, tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdministratorService_CreateEntity(t *testing.T) {
	type fields struct {
		usersService    usersService.UsersService
		cardService     CardService
		relationService RelationsServiceMethods
	}
	type args struct {
		ctx  context.Context
		user entity.User
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
			a := &AdministratorService{
				usersService:    tt.fields.usersService,
				cardService:     tt.fields.cardService,
				relationService: tt.fields.relationService,
			}
			if err := a.CreateEntity(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdministratorService_GetEntityReport(t *testing.T) {
	type fields struct {
		usersService    usersService.UsersService
		cardService     CardService
		relationService RelationsServiceMethods
	}
	type args struct {
		ctx      context.Context
		entityID int
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
			a := &AdministratorService{
				usersService:    tt.fields.usersService,
				cardService:     tt.fields.cardService,
				relationService: tt.fields.relationService,
			}
			if err := a.GetEntityReport(tt.args.ctx, tt.args.entityID); (err != nil) != tt.wantErr {
				t.Errorf("GetEntityReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
