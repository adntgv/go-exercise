package service

import (
	"reflect"
	"testing"
)

func Test_stringsToCurrencyPairs(t *testing.T) {
	type args struct {
		pairs []string
	}
	tests := []struct {
		name string
		args args
		want []currencyPair
	}{
		{
			args: args{
				pairs: []string{"BTC/USD", "BTC/CHF", "BTC/EUR"},
			},
			want: []currencyPair{"BTC/USD", "BTC/CHF", "BTC/EUR"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringsToCurrencyPairs(tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringsToCurrencyPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
