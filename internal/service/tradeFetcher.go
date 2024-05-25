package service

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/adntgv/go-exercise/pkg/kraken"
)

type TradeFetcher struct {
	client         *tradeClient
	pairs          []currencyPair
	ltpsMap        *sync.Map
	fetchingPeriod time.Duration
}

func newTradeFetcher(pairs []currencyPair, fetchingPeriosInSeconds int) *TradeFetcher {
	return &TradeFetcher{
		pairs: pairs,
		client: &tradeClient{
			fetcherFunc: kraken.GetLastTradedPrice,
		},
		ltpsMap:        &sync.Map{},
		fetchingPeriod: toFetchingPeriod(fetchingPeriosInSeconds),
	}
}

func toFetchingPeriod(seconds int) time.Duration {
	return time.Second * time.Duration(seconds)
}

func (s *TradeFetcher) GetLastTradedPrice() []LastTradedPrice {
	ltps := make([]LastTradedPrice, len(s.pairs))

	for i, pair := range s.pairs {
		value, ok := s.ltpsMap.Load(pair)
		if !ok {
			log.Println("could not get traded amount of ", pair)
			continue
		}

		ltp := LastTradedPrice{
			Pair:   pair,
			Amount: value.(amount),
		}

		ltps[i] = ltp
	}

	return ltps
}

func (s *TradeFetcher) Run() {
	for _, pair := range s.pairs {
		go func(pair currencyPair) {
			for {
				tradedAmount, err := s.fetchLatestTradedAmount(pair)
				if err != nil {
					log.Printf("could not fetch traded amount of '%v': %v", pair, err)
					continue
				}

				s.ltpsMap.Store(pair, tradedAmount)

				<-time.After(s.fetchingPeriod)
			}
		}(pair)
	}
}

func (s *TradeFetcher) fetchLatestTradedAmount(pair currencyPair) (amount, error) {
	lta, err := s.client.fetchLatestTradedAmount(pair)
	if err != nil {
		return "", fmt.Errorf("could not fetch latest traded amount: %v", err)
	}

	return lta, nil
}
