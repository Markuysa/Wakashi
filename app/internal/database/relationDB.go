package database

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
)

type RelationDB interface {
	BindEntities(ctx context.Context, masterUsername, slaveUsername string) error
}

func (db *BotDatabase) BindEntities(ctx context.Context, masterUsername, slaveUsername string) error {
	query := `
	insert into relation(
	                    master_id,
	                    slave_id
	)values (
	         $1,$2
	)
`
	master, err := db.GetUserID(ctx, masterUsername)
	if err != nil {
		return errors.New("cannot find the master")
	}
	slave, err := db.GetUserID(ctx, slaveUsername)
	if err != nil {
		return errors.New("cannot find the slave")
	}
	_, err = db.db.Query(ctx, query, master, slave)
	if err != nil {
		return errors.New("failed to bind entities:%v", err)
	}
	return nil
}
