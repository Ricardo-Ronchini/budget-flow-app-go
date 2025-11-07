package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/router"
)

func main() {
	log.Println("Servidor iniciado em http://localhost:8080")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("⚠️  Não foi possível carregar .env, usando variáveis do sistema, err: %v", err)
	}

	e := echo.New()
	router.Init(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servico de API. Err: %v", err)
	}
}
