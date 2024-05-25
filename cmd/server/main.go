package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adntgv/go-exercise/internal/service"
)

func main() {
	app := service.NewApp(currencies, fetchingPeriodInSeconds)

	go func() {
		app.Run()
	}()

	serve(port, app)
}

func serve(port string, app *service.App) {
	http.HandleFunc("/api/v1/ltp", app.Handle)

	serveAddress := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(serveAddress, nil); err != nil {
		log.Fatalln(err)
	}
}
