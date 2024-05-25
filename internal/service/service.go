package service

import (
	"sync"
)

type App struct {
	tradeFetcher *TradeFetcher
}

func NewApp(pairs []string) *App {
	return &App{
		tradeFetcher: newTradeFetcher(stringsToCurrencyPairs(pairs)),
	}
}

func stringsToCurrencyPairs(pairs []string) []currencyPair {
	cps := make([]currencyPair, len(pairs))

	for i, pair := range pairs {
		cps[i] = currencyPair(pair)
	}

	return cps
}

func (s *App) Run() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		s.tradeFetcher.Run()
		wg.Done()
	}()

	wg.Wait()
}

func (s *App) GetLastTradedPrices() []LastTradedPrice {
	return s.tradeFetcher.GetLastTradedPrice()
}
