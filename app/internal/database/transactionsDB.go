package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"tgBotIntern/app/internal/entity"
)

type TransactionDatabase interface {
	GetUsersTransactions(ctx context.Context, owerID int) ([]entity.Transaction, error)
}

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) GetUsersTransactions(ctx context.Context, owerID int) ([]entity.Transaction, error) {
	query := `
	select card_id,owner_id,operation_value,transaction_date from transactions 
	inner join users u on u.id = transactions.owner_id
	where u.id=$2
`
	rows, err := t.db.Query(ctx, query, owerID)
	if err != nil {
		return nil, errors.New("failed to get transactions")
	}
	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(&transaction)
		if err != nil {
			return nil, errors.New("failed to scan transaction")
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
