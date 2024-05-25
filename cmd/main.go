package main

import (
	"fmt"
	"net/http"

	"github.com/adntgv/go-exercise/internal/service"
	"github.com/labstack/echo/v4"
)

type response struct {
	Ltps []service.LastTradedPrice
}

func main() {
	app := service.NewApp([]string{"BTC/USD", "BTC/CHF", "BTC/EUR"})

	go func() {
		app.Run()
	}()

	e := echo.New()
	e.GET("/api/v1/ltp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Ltps: app.GetLastTradedPrices(),
		})
	})

	port := "80"
	serveAddress := fmt.Sprintf(":%v", port)
	e.Logger.Fatal(e.Start(serveAddress))
}
