package main

import (
	"os"
	"strings"
)

func init() {
	port = getEnv("PORT", port)
	currencies = getCurrenciesFromEnv(currencies)
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

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		return value
	}
}
