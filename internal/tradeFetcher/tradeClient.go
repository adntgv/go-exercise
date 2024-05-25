package tradefetcher

import (
	"fmt"

	"github.com/adntgv/go-exercise/internal/models"
)

type tradeClient struct {
	fetcherFunc func(string) (string, error)
}

func (c *tradeClient) fetchLatestTradedAmount(pair models.CurrencyPair) (models.Amount, error) {
	price, err := c.fetcherFunc(string(pair))

	if err != nil {
		return "", fmt.Errorf("could not fetch latest traded amount for %v: %v", pair, err)
	}

	return models.Amount(price), nil
}
