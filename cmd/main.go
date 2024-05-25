package main

import (
	"fmt"
	"net/http"

	"github.com/adntgv/go-exercise/internal/service"
	"github.com/labstack/echo/v4"
)

var (
	port       = "80"
	currencies = []string{"BTC/USD", "BTC/CHF", "BTC/EUR"}
)

func main() {
	app := service.NewApp(currencies)

	go func() {
		app.Run()
	}()

	serve(port, app)
}

func serve(port string, app *service.App) {
	e := echo.New()

	registerRoutes(e, app)

	serveAddress := fmt.Sprintf(":%v", port)
	e.Logger.Fatal(e.Start(serveAddress))
}

func registerRoutes(e *echo.Echo, app *service.App) {
	type response struct {
		Ltps []service.LastTradedPrice `json:"ltps"`
	}

	ltps := app.GetLastTradedPrices()

	e.GET("/api/v1/ltp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Ltps: ltps,
		})
	})
}
