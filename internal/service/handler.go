package service

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Ltps []LastTradedPrice
}

func (app *App) Handle(w http.ResponseWriter, r *http.Request) {
	resp := response{Ltps: app.GetLastTradedPrices()}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
