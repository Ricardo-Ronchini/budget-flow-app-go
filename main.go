package main

import (
	"log"
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/router"
)

func main() {
	r := router.Init() // Retorna o *http.ServeMux ou mux de algum framework

	log.Println("Servidor iniciado em http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error when start server:", err)
	}
}
