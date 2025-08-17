package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/router"
)

func main() {
	log.Println("Servidor iniciado em http://localhost:8080")

	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️  Não foi possível carregar .env, usando variáveis do sistema")
	}

	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
