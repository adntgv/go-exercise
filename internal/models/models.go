package models

type CurrencyPair string
type Amount string

type LastTradedPrice struct {
	Pair   CurrencyPair
	Amount Amount
}

type Response struct {
	Ltps []LastTradedPrice
}
