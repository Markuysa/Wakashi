package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"tgBotIntern/app/internal/entity"
)

// TransactionDatabase is an interface that provides contract to
// a structure that implements the functionality of a transaction table in the database
// The GetUsersTransactionsInWorkTime method is used to get transactions, that were created
// in as part of today's work shift
// The GetTurnover method is used to calculate the today's work shift turnover
// The AddTransactionRequest method is used to create new transaction request with unhandled status
// to delegate it to the collectors
// The GetAllUnhandledTransactions is used to get all transactions with status: unhandled
// The HandleTransaction method is used to mark the transaction as handled
// The GetTransactionByID method is used to get transaction data by id
type TransactionDatabase interface {
	GetUsersTransactionsInWorkTime(ctx context.Context, owerID int) ([]entity.Transaction, error)
	GetTurnover(ctx context.Context, ownerID int) (float64, error)
	AddTransactionRequest(ctx context.Context, transaction entity.Transaction) error
	GetAllUnhandledTransactions(ctx context.Context) ([]entity.Transaction, error)
	HandleTransaction(ctx context.Context, transactionID int) error
	GetTransactionByID(ctx context.Context, tranasctionID int) (entity.Transaction, error)
}

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}
func (t *TransactionRepository) GetTransactionByID(ctx context.Context, tranasctionID int) (entity.Transaction, error) {
	query := ` 
	select id,card_id,owner_id,operation_value,transaction_date,status,request_from from transactions
	where id=$1	
`
	var transaction entity.Transaction
	err := t.db.QueryRow(ctx, query, tranasctionID).Scan(
		&transaction.ID,
		&transaction.CardNumber,
		&transaction.OwnerID,
		&transaction.OperationValue,
		&transaction.TransactionDate,
		&transaction.Status,
		&transaction.RequestFromID,
	)
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}
func (t *TransactionRepository) HandleTransaction(ctx context.Context, transactionID int) error {
	query := `
	update transactions
	set status=true
	where id=$1
`
	_, err := t.db.Query(ctx, query, transactionID)
	if err != nil {
		return err
	}
	return nil
}
func (t *TransactionRepository) GetAllUnhandledTransactions(ctx context.Context) ([]entity.Transaction, error) {
	query := `
	select id,card_id,owner_id,operation_value,transaction_date,status,request_from from transactions
	where status=false	
`
	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, errors.New("failed to get transactions")
	}
	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.CardNumber,
			&transaction.OwnerID,
			&transaction.OperationValue,
			&transaction.TransactionDate,
			&transaction.Status,
			&transaction.RequestFromID,
		)
		if err != nil {
			return nil, errors.New("failed to scan transaction")
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
func (t *TransactionRepository) AddTransactionRequest(ctx context.Context, transaction entity.Transaction) error {
	query := `
	insert into transactions(
	                         card_id,
	                         owner_id,
	                         operation_value,
	                         transaction_date,
	                         status,
	                         request_from
	)values (
	         $1,$2,$3,$4,$5,$6
	);
`
	_, err := t.db.Query(ctx, query,
		transaction.CardNumber,
		transaction.OwnerID,
		transaction.OperationValue,
		transaction.TransactionDate,
		transaction.Status,
		transaction.RequestFromID,
	)
	if err != nil {
		return errors.New("failed to get transactions")
	}
	return nil
}
func (t *TransactionRepository) GetUsersTransactionsInWorkTime(ctx context.Context, owerID int) ([]entity.Transaction, error) {
	query := `
	select transactions.id,card_id,owner_id,operation_value,transaction_date,status, request_from from transactions
	inner join users u on u.id = transactions.owner_id
	where u.id=$1 and transaction_date BETWEEN current_date + time '08:00:00' AND CURRENT_DATE + time '12:00:00'
`
	rows, err := t.db.Query(ctx, query, owerID)
	if err != nil {
		return nil, errors.New("failed to get transactions")
	}
	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.CardNumber,
			&transaction.OwnerID,
			&transaction.OperationValue,
			&transaction.TransactionDate,
			&transaction.Status,
			&transaction.RequestFromID,
		)
		if err != nil {
			return nil, errors.New("failed to scan transaction")
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
func (t *TransactionRepository) GetTurnover(ctx context.Context, ownerID int) (float64, error) {

	query := `
	select SUM(operation_value) from transactions
	where owner_id=$1 and transaction_date BETWEEN current_date + time '08:00:00' AND CURRENT_DATE + time '12:00:00'
`
	var turover float64

	err := t.db.QueryRow(ctx, query, ownerID).Scan(&turover)
	if err != nil {
		return 0, errors.New("failed to calculate turnover")
	}
	return turover, nil
}
