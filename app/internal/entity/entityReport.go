package entity

type Report struct {
	Username     string
	Nickname     string
	RoleID       int
	Transactions []Transaction
}

func NewReport(username string, nickname string, roleID int, transactions []Transaction) *Report {
	return &Report{Username: username, Nickname: nickname, RoleID: roleID, Transactions: transactions}
}
