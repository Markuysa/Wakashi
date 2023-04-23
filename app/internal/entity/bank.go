package entity

type Bank struct {
	uin      int
	bankName string
}

func NewBank(uin int, bankName string) *Bank {
	return &Bank{uin: uin, bankName: bankName}
}
