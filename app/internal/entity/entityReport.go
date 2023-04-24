package entity

type Report struct {
	username     string
	roleID       int
	transactions []Transaction
}

func NewReport(username string, roleID int, transactions []Transaction) *Report {
	return &Report{username: username, roleID: roleID, transactions: transactions}
}
