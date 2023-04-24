package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/hedzr/errors.v3"
)

type RelationDatabase interface {
	BindEntities(ctx context.Context, master, slave int) error
}
type RelationRepository struct {
	db *pgxpool.Pool
}

func NewRelationDB(db *pgxpool.Pool) *RelationRepository {
	return &RelationRepository{db: db}
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
