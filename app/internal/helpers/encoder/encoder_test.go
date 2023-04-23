package encoder

import "testing"

func TestIsMatch(t *testing.T) {
	password, _ := EncodePassword("islam20011")
	type args struct {
		encoded  string
		original string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "correctPassword",
			args: args{original: "islam20011", encoded: password}, want: true, wantErr: false},
		{name: "incorrectPassword",
			args: args{original: "islam2001", encoded: password}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsMatch(tt.args.encoded, tt.args.original)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
