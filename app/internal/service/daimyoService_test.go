package service

import (
	"context"
	"reflect"
	"testing"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

func TestDaimyoService_BindShogun(t *testing.T) {
	type fields struct {
		cardsService     CardRights
		userService      usersService.UsersRepositoryService
		relationsService RelationsServiceMethods
	}
	type args struct {
		ctx            context.Context
		shogunUsername string
		daimyoUsername string
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
			s := &DaimyoService{
				cardsService:     tt.fields.cardsService,
				userService:      tt.fields.userService,
				relationsService: tt.fields.relationsService,
			}
			if err := s.BindShogun(tt.args.ctx, tt.args.shogunUsername, tt.args.daimyoUsername); (err != nil) != tt.wantErr {
				t.Errorf("BindShogun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaimyoService_CreateCardIncreasementRequest(t *testing.T) {
	type fields struct {
		cardsService     CardRights
		userService      usersService.UsersRepositoryService
		relationsService RelationsServiceMethods
	}
	type args struct {
		ctx               context.Context
		cardID            int
		increasementValue float64
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
			s := &DaimyoService{
				cardsService:     tt.fields.cardsService,
				userService:      tt.fields.userService,
				relationsService: tt.fields.relationsService,
			}
			if err := s.CreateCardIncreasementRequest(tt.args.ctx, tt.args.cardID, tt.args.increasementValue); (err != nil) != tt.wantErr {
				t.Errorf("CreateCardIncreasementRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaimyoService_GetCardsList(t *testing.T) {
	type fields struct {
		cardsService     CardRights
		userService      usersService.UsersRepositoryService
		relationsService RelationsServiceMethods
	}
	type args struct {
		ctx     context.Context
		ownerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Card
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DaimyoService{
				cardsService:     tt.fields.cardsService,
				userService:      tt.fields.userService,
				relationsService: tt.fields.relationsService,
			}
			got, err := s.GetCardsList(tt.args.ctx, tt.args.ownerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCardsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCardsList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaimyoService_GetSamuraiList(t *testing.T) {
	type fields struct {
		cardsService     CardRights
		userService      usersService.UsersRepositoryService
		relationsService RelationsServiceMethods
	}
	type args struct {
		ctx            context.Context
		masterUsername string
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
			s := &DaimyoService{
				cardsService:     tt.fields.cardsService,
				userService:      tt.fields.userService,
				relationsService: tt.fields.relationsService,
			}
			got, err := s.GetSamuraiList(tt.args.ctx, tt.args.masterUsername)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSamuraiList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSamuraiList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaimyoService_SetCardsBalances(t *testing.T) {
	type fields struct {
		cardsService     CardRights
		userService      usersService.UsersRepositoryService
		relationsService RelationsServiceMethods
	}
	type args struct {
		ctx        context.Context
		totalValue float64
		cardnumber int
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
			s := &DaimyoService{
				cardsService:     tt.fields.cardsService,
				userService:      tt.fields.userService,
				relationsService: tt.fields.relationsService,
			}
			if err := s.SetCardsBalances(tt.args.ctx, tt.args.totalValue, tt.args.cardnumber); (err != nil) != tt.wantErr {
				t.Errorf("SetCardsBalances() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
