package service

import (
	"github.com/adntgv/go-exercise/internal/models"
	tradefetcher "github.com/adntgv/go-exercise/internal/tradeFetcher"
)

type App struct {
	tradeFetcher *tradefetcher.TradeFetcher
}

func NewApp(pairs []string, fetchingPeriodInSeconds int) *App {
	return &App{
		tradeFetcher: tradefetcher.New(stringsToCurrencyPairs(pairs), fetchingPeriodInSeconds),
	}
}

func stringsToCurrencyPairs(pairs []string) []models.CurrencyPair {
	cps := make([]models.CurrencyPair, len(pairs))

	for i, pair := range pairs {
		cps[i] = models.CurrencyPair(pair)
	}

	return cps
}

func (s *App) Run() {
	s.tradeFetcher.Run()
}

func (s *App) GetLastTradedPrices() []models.LastTradedPrice {
	return s.tradeFetcher.GetLastTradedPrice()
}
