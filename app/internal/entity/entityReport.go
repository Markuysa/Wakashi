package entity

type Report struct {
	username string
	roleID   int
	turnover float64
}

func NewReport(username string, roleID int, turnover float64) *Report {
	return &Report{username: username, roleID: roleID, turnover: turnover}
}
