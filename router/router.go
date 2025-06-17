package router

import (
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/auth"
	"github.com/ricardo-ronchini/budget-flow-app-go/handler"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	// Rotas públicas
	mux.HandleFunc("/login", handler.Login)

	// Rotas protegidas
	mux.Handle("/expenses", auth.Middleware(http.HandlerFunc(handler.Expenses)))
	// mux.Handle("/gastos/novo", auth.Middleware(http.HandlerFunc(handler.CriarGasto)))

	return mux
}
