package models

// AdministratorRights is an interface which represents
// the methods available to the administrator
// CreateEntity method crates new entity object
// CreateCard method created new bank card
// BindSlave method binds slave object with master object
// BindCardToDaimyo method binds a new card to Daimyo
// GetEntityReport method creates full information about any entity
type AdministratorRights interface {
	CreateEntity(entityName string) (Entity, error)
	CreateCard() (*Card, error)
	BindSlave() error
	BindCardToDaimyo(daimyoId int) error
	GetEntityReport(entity Entity) error
}
