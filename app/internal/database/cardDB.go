package database

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/entity"
)

type CardsDB interface {
	GetCard(ctx context.Context, cardNumber int) (*entity.Card, error)
	AddCard(ctx context.Context, card entity.Card) error
	BindCard(ctx context.Context, cardNumber int, ownerID int) error
	GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error)
	SetCardTotal(ctx context.Context, total float64, number int) error
	IncreaseTotal(ctx context.Context, incValue float64) error
	CalculateTurnover(ctx context.Context, username string) (float64, error)
}

func (db *BotDatabase) CalculateTurnover(ctx context.Context, username string) (float64, error) {
	query := `
		select SUM(total) from card inner join users u on u.id = card.owner_id
		where username=$1
		
`
	var total float64
	err := db.db.QueryRow(ctx, query, username).Scan(&total)
	if err != nil {
		return 0, errors.New("failed to get total")
	}
	return total, nil
}
func (db *BotDatabase) IncreaseTotal(ctx context.Context, incValue float64) error {

	return nil
}
func (db *BotDatabase) SetCardTotal(ctx context.Context, total float64, number int) error {
	query := `
		update card
		set total=$1
		where card_number=$2
	`
	_, err := db.db.Query(ctx, query, total, number)
	return err
}
func (db *BotDatabase) GetCard(ctx context.Context, cardNumber int) (*entity.Card, error) {

	query := `
	select bank_id,card_number,card_limit,owner_id,cvv_code from card
	where card_number=$1
`
	var card entity.Card
	err := db.db.QueryRow(ctx, query, cardNumber).Scan(&card)
	if err != nil {
		return nil, errors.New("failed to get card info:%v", err)
	}
	return &card, nil
}
func (db *BotDatabase) AddCard(ctx context.Context, card entity.Card) error {

	query := `
	insert into card(
					 bank_id,
					 card_number,
	                 card_limit,
	                 owner_id,
	                 cvv_code
	) values (
	        $1,$2,$3,$4,$5
	)
`
	_, err := db.db.Query(ctx, query, card.IssuerBankID,
		card.CardNumber,
		card.DailyLimit,
		card.DaimyoID,
		card.CvvCode,
	)
	if err != nil {
		return errors.New("failed to add card:%v", err)
	}
	return nil
}
func (db *BotDatabase) BindCard(ctx context.Context, cardNumber int, ownerID int) error {
	_, err := db.GetCard(ctx, cardNumber)
	if err != nil {
		return errors.New("card doesn't exist")
	}
	query := `
	update card
	set owner_id=$1
	where card_number=$2;
`
	_, err = db.db.Query(ctx, query, ownerID, cardNumber)
	if err != nil {
		return errors.New("failed to bind card: %v", err)
	}
	return nil
}
func (db *BotDatabase) GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error) {
	query := `
	select * from card
	where owner_id=$1
`
	rows, err := db.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, errors.New("failed to get cards list:%v", err)
	}
	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		err := rows.Scan(&card)
		if err != nil {
			return nil, errors.New("failed to scan cards list:%v", err)
		}
		cards = append(cards, card)
	}
	return cards, nil
}
