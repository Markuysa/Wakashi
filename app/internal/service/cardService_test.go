package service

import (
	"context"
	"reflect"
	"testing"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
)

func TestCardService_BindToDaimyo(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
	}
	type args struct {
		ctx        context.Context
		cardNumber int
		ownerID    int
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
			s := &CardService{
				cardsRepos: tt.fields.cardsRepos,
			}
			if err := s.BindToDaimyo(tt.args.ctx, tt.args.cardNumber, tt.args.ownerID); (err != nil) != tt.wantErr {
				t.Errorf("BindToDaimyo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCardService_CreateCard(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
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
			s := &CardService{
				cardsRepos: tt.fields.cardsRepos,
			}
			if err := s.CreateCard(tt.args.ctx, tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("CreateCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCardService_GetCardsList(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
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
			s := &CardService{
				cardsRepos: tt.fields.cardsRepos,
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

func TestCardService_GetTurnover(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
	}
	type args struct {
		ctx            context.Context
		daimyoUsername string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CardService{
				cardsRepos: tt.fields.cardsRepos,
			}
			got, err := c.GetTurnover(tt.args.ctx, tt.args.daimyoUsername)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTurnover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTurnover() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardService_HandleCardTotalInc(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
	}
	type args struct {
		ctx      context.Context
		incValue float64
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
			c := &CardService{
				cardsRepos: tt.fields.cardsRepos,
			}
			if err := c.HandleCardTotalInc(tt.args.ctx, tt.args.incValue); (err != nil) != tt.wantErr {
				t.Errorf("HandleCardTotalInc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCardService_SetTotal(t *testing.T) {
	type fields struct {
		cardsRepos database.CardsDB
	}
	type args struct {
		ctx        context.Context
		total      float64
		cardNumber int
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
			s := &CardService{
				cardsRepos: tt.fields.cardsRepos,
			}
			if err := s.SetTotal(tt.args.ctx, tt.args.total, tt.args.cardNumber); (err != nil) != tt.wantErr {
				t.Errorf("SetTotal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
