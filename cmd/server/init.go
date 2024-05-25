package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	port                    = "80"
	currencies              = []string{"BTC/USD", "BTC/CHF", "BTC/EUR"}
	fetchingPeriodInSeconds = 30
)

func init() {
	port = getEnv("PORT", port)
	currencies = getCurrenciesFromEnv(currencies)
	fetchingPeriodInSeconds = getFetchingPeriod(fetchingPeriodInSeconds)
}

func getCurrenciesFromEnv(defaultValue []string) []string {
	if value, ok := os.LookupEnv("CURRENCIES"); !ok {
		return defaultValue
	} else {
		pairs := strings.Split(value, ",")

		for i, pair := range pairs {
			pairs[i] = strings.TrimSpace(pair)
		}

		return pairs
	}
}

func getFetchingPeriod(defaultValue int) int {
	if value, ok := os.LookupEnv("FETCHING_PERIOD_IN_SECONDS"); !ok {
		return defaultValue
	} else {
		result, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return result
	}
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		return value
	}
}
