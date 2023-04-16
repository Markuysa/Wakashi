package models

type SamuraiRights interface {
	SetTurnover() error
	BindToDamiyo() error
}
