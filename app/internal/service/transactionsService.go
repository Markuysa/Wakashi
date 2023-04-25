package service

import (
	"context"
	"errors"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

type TransactionProcessor interface {
	GetTransactions(ctx context.Context, username string) ([]entity.Transaction, error)
	GetUnhandledTransactions(ctx context.Context) ([]entity.Transaction, error)
	GetTurnover(ctx context.Context, username string) (float64, error)
	CreateTransactionRequest(ctx context.Context, transaction entity.Transaction) error
	HandleTransaction(ctx context.Context, transactionID int) error
	GetTransaction(ctx context.Context, transactionID int) (entity.Transaction, error)
}
type TransactionsService struct {
	transactionDB database.TransactionDatabase
	usersService  usersService.UsersRepositoryService
}

func NewTransactionsService(transactionDB database.TransactionDatabase, usersService usersService.UsersRepositoryService) *TransactionsService {
	return &TransactionsService{transactionDB: transactionDB, usersService: usersService}
}
func (t *TransactionsService) GetTransaction(ctx context.Context, transactionID int) (entity.Transaction, error) {
	return t.transactionDB.GetTransactionByID(ctx, transactionID)
}
func (t *TransactionsService) CreateTransactionRequest(ctx context.Context, transaction entity.Transaction) error {
	return t.transactionDB.AddTransactionRequest(ctx, transaction)
}
func (t *TransactionsService) HandleTransaction(ctx context.Context, transactionID int) error {
	return t.transactionDB.HandleTransaction(ctx, transactionID)
}
func (t *TransactionsService) GetUnhandledTransactions(ctx context.Context) ([]entity.Transaction, error) {
	return t.transactionDB.GetAllUnhandledTransactions(ctx)
}
func (t *TransactionsService) GetTransactions(ctx context.Context, username string) ([]entity.Transaction, error) {
	userID, err := t.usersService.GetUserID(ctx, username)
	if err != nil {
		return nil, errors.New("No user with username: " + username)
	}
	return t.transactionDB.GetUsersTransactionsInWorkTime(ctx, userID)
}
func (t *TransactionsService) GetTurnover(ctx context.Context, username string) (float64, error) {
	userID, err := t.usersService.GetUserID(ctx, username)
	if err != nil {
		return 0, errors.New("No user with username: " + username)
	}
	return t.transactionDB.GetTurnover(ctx, userID)
}
