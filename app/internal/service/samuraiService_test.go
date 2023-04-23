package service

import "testing"

func TestSamuraiService_BindToDamiyo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamuraiService{}
			if err := s.BindToDamiyo(); (err != nil) != tt.wantErr {
				t.Errorf("BindToDamiyo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSamuraiService_SetTurnover(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SamuraiService{}
			if err := s.SetTurnover(); (err != nil) != tt.wantErr {
				t.Errorf("SetTurnover() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
