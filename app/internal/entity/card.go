package entity

type Card struct {
	DaimyoID     int
	IssuerBankID int
	CvvCode      int
	CardNumber   string
	DailyLimit   int
	Total        float64
}

func NewCard(daimyoID int, issuerBankID int, cvvCode int, cardNumber string, dailyLimit int, total float64) *Card {
	return &Card{DaimyoID: daimyoID,
		IssuerBankID: issuerBankID,
		CvvCode:      cvvCode,
		CardNumber:   cardNumber,
		DailyLimit:   dailyLimit,
		Total:        total}
}
