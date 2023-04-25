package service

import (
	"context"
	"errors"
)

// CollectorRights is an interface that provides contract for CollectorService
// struct represents one of the parts of use-case layer
// The HandleDaimyoIncreasementRequest method is used to handle the daimyo request
// and change the card balance
type CollectorRights interface {
	HandleDaimyoIncreasementRequest(ctx context.Context, requestID int) error
}

type CollectorService struct {
	cardService        CardRights
	transactionService TransactionProcessor
}

func NewCollectorService(cardService CardRights, transactionService TransactionProcessor) *CollectorService {
	return &CollectorService{cardService: cardService, transactionService: transactionService}
}
func (c *CollectorService) HandleDaimyoIncreasementRequest(ctx context.Context, requestID int) error {
	transaction, err := c.transactionService.GetTransaction(ctx, requestID)
	if err != nil {
		return errors.New("cannot get transaction value")
	}
	err = c.cardService.HandleCardTotalInc(ctx, transaction.OperationValue, transaction.CardNumber)
	if err != nil {
		return errors.New("cannot update card balance")
	}
	return c.transactionService.HandleTransaction(ctx, requestID)
}
