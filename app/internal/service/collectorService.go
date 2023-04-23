package service

import "context"

type CollectorRights interface {
	HandleDaimyoIncreasementRequest(ctx context.Context, incValue float64) error
}

type CollectorService struct {
	cardService CardService
}

func (c *CollectorService) HandleDaimyoIncreasementRequest(ctx context.Context, incValue float64) error {
	return c.cardService.HandleCardTotalInc(ctx, incValue)
}
