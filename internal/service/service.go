package service

type App struct {
	tradeFetcher *TradeFetcher
}

func NewApp(pairs []string, fetchingPeriodInSeconds int) *App {
	return &App{
		tradeFetcher: newTradeFetcher(stringsToCurrencyPairs(pairs), fetchingPeriodInSeconds),
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
	s.tradeFetcher.Run()
}

func (s *App) GetLastTradedPrices() []LastTradedPrice {
	return s.tradeFetcher.GetLastTradedPrice()
}
