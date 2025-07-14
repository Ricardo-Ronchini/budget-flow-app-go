package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/router"
)

func main() {
	log.Println("Servidor iniciado em http://localhost:9000")

	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":8080"))

	// if err := http.ListenAndServe(":9000", router.Init()); err != nil {
	// 	log.Fatal("Error when start server:", err)
	// }
}
