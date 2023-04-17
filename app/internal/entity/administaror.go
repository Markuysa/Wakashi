package entity

import "tgBotIntern/app/internal/constants/roles"

type Administrator struct {
	id int
}

// AdministratorRights is an interface which represents
// the methods available to the administrator
// CreateEntity method crates new entity object
// CreateCard method created new bank card
// BindSlave method binds slave object with master object
// BindCardToDaimyo method binds a new card to Daimyo
// GetEntityReport method creates full information about any entity
type AdministratorRights interface {
	CreateEntity(entityId int) (Entity, error)
	CreateCard() (*Card, error)
	BindSlave() error
	BindCardToDaimyo(daimyoId int) error
	GetEntityReport(entity Entity) error
}

func (a *Administrator) CreateEntity(entityId int) (Entity, error) {
	switch entityId {
	case roles.Administrator:
		{
			// add service to communicate with database
			return Administrator{}, nil
		}
	case roles.Shogun:
		{
			return Shogun{}, nil
		}
	case roles.Daimyo:
		{
			return Shogun{}, nil
		}
	case roles.Samurai:
		{
			return Shogun{}, nil
		}
	case roles.Collector:
		{
			return Shogun{}, nil
		}
	case roles.Card:
		{
			return Card{}, nil
		}
	}
	return nil, nil
}

func (a *Administrator) CreateCard() (*Card, error) {
	return nil, nil
}
func (a *Administrator) BindSlave() error {
	return nil
}

func (a *Administrator) BindCardToDaimyo(daimyoId int) error {
	return nil
}
func (a *Administrator) GetEntityReport(entity Entity) error {
	return nil
}
