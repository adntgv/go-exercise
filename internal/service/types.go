package service

type currencyPair string
type amount string

type LastTradedPrice struct {
	Pair   currencyPair
	Amount amount
}
