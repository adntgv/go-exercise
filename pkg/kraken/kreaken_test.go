package kraken

import "testing"

func TestGetLastTradedPrice(t *testing.T) {
	type args struct {
		pair string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple get",
			args: args{
				pair: "BTC/CHF",
			},
			want:    "63011.50000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLastTradedPrice(tt.args.pair)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastTradedPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLastTradedPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
