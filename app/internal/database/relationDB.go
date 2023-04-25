package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/hedzr/errors.v3"
)

// RelationDatabase is an interface that provides contract to relation repository objects
// The BindEntities method is used to bind two entities with each other
// The GetMasterUsername method is used to get the master of user with given ID
type RelationDatabase interface {
	BindEntities(ctx context.Context, master, slave int) error
	GetMasterUsername(ctx context.Context, slaveID int) (string, error)
}
type RelationRepository struct {
	db *pgxpool.Pool
}

func NewRelationDB(db *pgxpool.Pool) *RelationRepository {
	return &RelationRepository{db: db}
}
func (db *RelationRepository) GetMasterUsername(ctx context.Context, slaveID int) (string, error) {
	query := `
	select username from users inner join relation r on users.id=r.master_id
	where r.slave_id=$1
`
	var username string
	err := db.db.QueryRow(ctx, query, slaveID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
func (db *RelationRepository) BindEntities(ctx context.Context, master, slave int) error {
	query := `
	insert into relation(
	                    master_id,
	                    slave_id
	)values (
	         $1,$2
	)
`
	_, err := db.db.Query(ctx, query, master, slave)
	if err != nil {
		return errors.New("failed to bind entities:%v", err)
	}
	return nil
}
