package service

import (
	"context"
	"testing"
)

func TestCollectorService_HandleDaimyoIncreasementRequest(t *testing.T) {
	type fields struct {
		cardService CardService
	}
	type args struct {
		ctx      context.Context
		incValue float64
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
			c := &CollectorService{
				cardService: tt.fields.cardService,
			}
			if err := c.HandleDaimyoIncreasementRequest(tt.args.ctx, tt.args.incValue); (err != nil) != tt.wantErr {
				t.Errorf("HandleDaimyoIncreasementRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
