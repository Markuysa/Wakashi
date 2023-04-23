package service

import (
	"context"
	"testing"
	"tgBotIntern/app/internal/database"
)

func TestRelationsService_Bind(t *testing.T) {
	type fields struct {
		relationDB database.RelationDB
	}
	type args struct {
		ctx            context.Context
		masterUsername string
		slaveUsername  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RelationsService{
				relationDB: tt.fields.relationDB,
			}
			if err := s.Bind(tt.args.ctx, tt.args.masterUsername, tt.args.slaveUsername); (err != nil) != tt.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
