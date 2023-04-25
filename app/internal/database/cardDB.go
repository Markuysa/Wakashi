package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/entity"
)

type CardsDatabase interface {
	GetCard(ctx context.Context, cardNumber int) (*entity.Card, error)
	AddCard(ctx context.Context, card entity.Card) error
	BindCard(ctx context.Context, cardNumber int, ownerID int) error
	GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error)
	SetCardTotal(ctx context.Context, total float64, number int) error
	IncreaseTotal(ctx context.Context, incValue float64, cardNumber int) error
	CalculateTurnover(ctx context.Context, userID, requestFromID int) (float64, error)
	GetCardsTotal(ctx context.Context, username string) (float64, error)
}
type CardsRepository struct {
	db *pgxpool.Pool
}

func NewCardsDB(db *pgxpool.Pool) *CardsRepository {
	return &CardsRepository{db: db}
}
func (db *CardsRepository) GetCardsTotal(ctx context.Context, username string) (float64, error) {
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
func (db *CardsRepository) CalculateTurnover(ctx context.Context, userID, requestFromID int) (float64, error) {
	query := `
		select SUM(operation_value)  from transactions inner join users u on transactions.owner_id = u.id
		where u.id=$1 and request_from=$2  and transaction_date BETWEEN current_date + time '08:00:00' AND CURRENT_DATE + time '12:00:00'
		
`
	var total float64
	err := db.db.QueryRow(ctx, query, userID, requestFromID).Scan(&total)
	if err != nil {
		return 0, errors.New("failed to get total")
	}
	return total, nil
}
func (db *CardsRepository) IncreaseTotal(ctx context.Context, incValue float64, cardNumber int) error {

	query := `
	update card
	set total = total + $1
	where card_number=$2
`
	_, err := db.db.Query(ctx, query, incValue, cardNumber)
	if err != nil {
		return errors.New("failed to update card number")
	}
	return nil
}
func (db *CardsRepository) SetCardTotal(ctx context.Context, total float64, number int) error {
	query := `
		update card
		set total=$1
		where card_number=$2
	`
	_, err := db.db.Query(ctx, query, total, number)
	return err
}
func (db *CardsRepository) GetCard(ctx context.Context, cardNumber int) (*entity.Card, error) {

	query := `
	select bank_id,card_number,card_limit,owner_id,cvv_code from card
	where card_number=$1
`
	var card entity.Card
	err := db.db.QueryRow(ctx, query, cardNumber).Scan(&card.IssuerBankID, &card.CardNumber,
		&card.DailyLimit, &card.DaimyoID, &card.CvvCode)
	if err != nil {
		return nil, errors.New("failed to get card info:%v", err)
	}
	return &card, nil
}
func (db *CardsRepository) AddCard(ctx context.Context, card entity.Card) error {

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
func (db *CardsRepository) BindCard(ctx context.Context, cardNumber int, ownerID int) error {
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
func (db *CardsRepository) GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error) {
	query := `
	select owner_id,
	       card_number,
	       bank_id,
	       card_limit,
	       cvv_code,
	       total
	from card
	where owner_id=$1
`
	rows, err := db.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, errors.New("failed to get cards list:%v", err)
	}
	var cards []entity.Card
	for rows.Next() {
		var card entity.Card
		err := rows.Scan(&card.DaimyoID,
			&card.CardNumber, &card.IssuerBankID,
			&card.DailyLimit, &card.CvvCode,
			&card.Total)
		if err != nil {
			return nil, errors.New("failed to scan cards list:%v", err)
		}
		cards = append(cards, card)
	}
	return cards, nil
}
