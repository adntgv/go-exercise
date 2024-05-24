package kraken

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const host = "https://api.kraken.com/0/public/Ticker?pair="

// Define the structs matching the JSON structure
type Response struct {
	Error  []interface{}         `json:"error"`
	Result map[string]TickerInfo `json:"result"`
}

type TickerInfo struct {
	A [3]string `json:"a"`
	B [3]string `json:"b"`
	C [2]string `json:"c"`
	V [2]string `json:"v"`
	P [2]string `json:"p"`
	T [2]int    `json:"t"`
	L [2]string `json:"l"`
	H [2]string `json:"h"`
	O string    `json:"o"`
}

func GetLastTradedPrice(pair string) (string, error) {
	url := fmt.Sprintf("%v%v", host, pair)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get %v: %v", url, err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	result := new(Response)
	err = json.Unmarshal(bz, result)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if len(result.Error) > 0 {
		return "", fmt.Errorf("responded with error: %v", result.Error)
	}

	item, ok := result.Result[pair]
	if !ok {
		return "", fmt.Errorf("response does not contain target currency pair %v: %v", pair, result)
	}

	price := item.C[0]

	return price, nil
}
