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
}
type TransactionsService struct {
	transactionDB database.TransactionDatabase
	usersService  usersService.UsersRepositoryService
}

func NewTransactionsService(transactionDB database.TransactionDatabase, usersService usersService.UsersRepositoryService) *TransactionsService {
	return &TransactionsService{transactionDB: transactionDB, usersService: usersService}
}

func (t *TransactionsService) GetTransactions(ctx context.Context, username string) ([]entity.Transaction, error) {
	userID, err := t.usersService.GetUserID(ctx, username)
	if err != nil {
		return nil, errors.New("No user with username: " + username)
	}
	return t.transactionDB.GetUsersTransactions(ctx, userID)
}
