package entity

type Card struct {
	DaimyoID     int
	IssuerBankID int
	CvvCode      string
	CardNumber   int64
	DailyLimit   int
	Total        float64
}

func NewCard(daimyoID int, issuerBankID int, cvvCode string, cardNumber int64, dailyLimit int, total float64) *Card {
	return &Card{DaimyoID: daimyoID,
		IssuerBankID: issuerBankID,
		CvvCode:      cvvCode,
		CardNumber:   cardNumber,
		DailyLimit:   dailyLimit,
		Total:        total}
}
