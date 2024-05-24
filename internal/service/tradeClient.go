package service

import "fmt"

type tradeClient struct {
	fetcherFunc func(string) (string, error)
}

func (c *tradeClient) fetchLatestTradedAmount(pair currencyPair) (amount, error) {
	price, err := c.fetcherFunc(string(pair))

	if err != nil {
		return "", fmt.Errorf("could not fetch latest traded amount for %v: %v", pair, err)
	}

	return amount(price), nil
}
