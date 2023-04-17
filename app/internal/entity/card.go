package entity

type CardRights interface {
	BindToDaimyo(daimyoId int)
}

type Card struct {
	BankData *Bank
}
