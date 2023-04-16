package models

type CollectorRights interface {
	HandleDaimyoIncreasementRequest() error
}
