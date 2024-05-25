package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adntgv/go-exercise/internal/models"
	"github.com/adntgv/go-exercise/internal/service"
)

func TestGetLTPHandler(t *testing.T) {
	currencies := []string{"BTC/USD"}
	fetchingPeriod := 1

	app := service.NewApp(currencies, fetchingPeriod)

	go func() {
		app.Run()
	}()

	<-time.After(time.Second)

	handler := http.HandlerFunc(app.Handle)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/ltp", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}

	bz := rr.Body.Bytes()

	response := models.Response{}

	if err = json.Unmarshal(bz, &response); err != nil {
		t.Errorf("handler returned invalid response: got %v", err)
		return
	}

	if len(response.Ltps) < 1 {
		t.Errorf("handler returned empty list: got %v want %v", 0, 1)
		return
	}

	for _, currency := range currencies {
		found := false
		for _, item := range response.Ltps {
			if string(item.Pair) == currency {
				found = true
			}
		}

		if !found {
			t.Errorf("did not find currency %v in %v", currency, response.Ltps)
			return
		}
	}
}
