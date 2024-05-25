package kraken

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetLastTradedPrice(t *testing.T) {
	type args struct {
		pair string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple get",
			args: args{
				pair: "BTC/CHF",
			},
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

			value, err := strconv.ParseFloat(got, 64)
			if err != nil {
				t.Errorf("GetLastTradedPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if value == 0 {
				t.Errorf("GetLastTradedPrice() error = %v, wantErr %v", fmt.Errorf("value is equal to 0"), tt.wantErr)
				return
			}
		})
	}
}
