package entity

type Report struct {
	Username     string
	RoleID       int
	Transactions []Transaction
}

func NewReport(username string, roleID int, transactions []Transaction) *Report {
	return &Report{Username: username, RoleID: roleID, Transactions: transactions}
}
