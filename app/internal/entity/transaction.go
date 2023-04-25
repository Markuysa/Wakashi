package entity

import "time"

type Transaction struct {
	ID              int
	RequestFromID   int
	CardNumber      int
	OwnerID         int
	OperationValue  float64
	TransactionDate time.Time
	Status          bool
}
