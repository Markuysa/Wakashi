package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tgBotIntern/app/internal/database"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/pkg/auth/service/usersService"
)

// CardRights is an interface that provides contract to CardService struct
// that represents use-case layer to interact with bank cards
// The BindToDaimyo method is used to bind card to daimyo
// The CreateCard method is used to create new bank card
// The GetCardsList method is used to get cards list of username with given id
// The SetTotal method is used to set card balance
// The HandleCardTotalInc method is used to update card balance after completing transaction
// The GetTurnover method is used to get turnover of the card with given owner and request dealer id's
// Which means that not only the owner of card (in out case daimyo) can create requests to replenish the balance
// Also samurai have that right, so we need to use requestDealer username to get access to person, who
// created the replenishment request
// The GetCard method is used to get card by its number
// The GetTotal method is used to get balance of the user card with given username
type CardRights interface {
	BindToDaimyo(ctx context.Context, cardNumber, ownerUsername string) error
	CreateCard(ctx context.Context, card entity.Card) error
	GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error)
	SetTotal(ctx context.Context, total float64, cardNumber int) error
	HandleCardTotalInc(ctx context.Context, incValue float64, cardNumber int) error
	GetTurnover(ctx context.Context, ownerUsername, requestDealerUsername string) (float64, error)
	GetCard(ctx context.Context, cardNumber int) (*entity.Card, error)
	GetTotal(ctx context.Context, username string) (float64, error)
}

type CardService struct {
	cardsRepos  database.CardsDatabase
	userService usersService.UsersRepositoryService
}

func NewCardService(cardsRepos database.CardsDatabase, userService usersService.UsersRepositoryService) *CardService {
	return &CardService{cardsRepos: cardsRepos, userService: userService}
}
func (c *CardService) GetCard(ctx context.Context, cardNumber int) (*entity.Card, error) {
	return c.cardsRepos.GetCard(ctx, cardNumber)
}
func (c *CardService) GetTotal(ctx context.Context, username string) (float64, error) {
	return c.cardsRepos.GetCardsTotal(ctx, username)
}
func (c *CardService) GetTurnover(ctx context.Context, ownerUsername, requestDealerUsername string) (float64, error) {
	owner, err := c.userService.GetUserID(ctx, ownerUsername)
	if err != nil {
		return 0, errors.New("no user with username: " + ownerUsername)
	}
	dealer, err := c.userService.GetUserID(ctx, requestDealerUsername)
	if err != nil {
		return 0, errors.New("no user with username: " + requestDealerUsername)
	}
	return c.cardsRepos.CalculateTurnover(ctx, owner, dealer)
}

func (c *CardService) HandleCardTotalInc(ctx context.Context, incValue float64, cardNumber int) error {
	return c.cardsRepos.IncreaseTotal(ctx, incValue, cardNumber)
}
func (s *CardService) CreateCard(ctx context.Context, card entity.Card) error {
	return s.cardsRepos.AddCard(ctx, card)
}
func (s *CardService) BindToDaimyo(ctx context.Context, cardNumber, ownerUsername string) error {
	ownerID, err := s.userService.GetUserID(ctx, ownerUsername)
	if err != nil {
		return errors.New("No user with that username: " + ownerUsername)
	}
	number, err := strconv.Atoi(cardNumber)
	if err != nil {
		return errors.New("Check correctness of your card number!")
	}
	return s.cardsRepos.BindCard(ctx, number, ownerID)
}
func (s *CardService) GetCardsList(ctx context.Context, ownerID int) ([]entity.Card, error) {
	return s.cardsRepos.GetCardsList(ctx, ownerID)
}
func (s *CardService) SetTotal(ctx context.Context, total float64, cardNumber int) error {
	return s.cardsRepos.SetCardTotal(ctx, total, cardNumber)
}
