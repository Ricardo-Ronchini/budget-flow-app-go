package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/auth"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/handler"
)

func Init(e *echo.Echo) {
	// Estudar tecnologia das rotas a ser implementada
	// *** trocar para gorila mux? - Esta sem manutençao faz tempo
	// Gin?
	// Chi?
	// Echo?
	// mux := http.NewServeMux()

	// Rotas públicas
	// mux.HandleFunc("/login", handler.Login) // precisa de ajuste

	// Rotas protegidas
	// mux.Handle("/expenses", auth.Middleware(http.HandlerFunc(handler.Expenses)))
	// mux.Handle("/gastos/novo", auth.Middleware(http.HandlerFunc(handler.CriarGasto)))

	// *** Echo ***

	e.POST("login", handler.Login)

	// add middleware
	api := e.Group("/api", auth.Middleware)

	routes := []*contexts.WebRoute{
		handler.V1ExpensesGET,
		handler.V1ExpensesPOST,
		handler.V1ExpensesPUT,
		handler.V1ExpensesPATCH,
		handler.V1ExpensesDELETE,
	}

	for _, r := range routes {
		register(api, r)
	}
}

func register(g *echo.Group, wr *contexts.WebRoute) {
	switch wr.Method {
	case "GET":
		g.GET(wr.Path, wr.Handler)
	case "POST":
		g.POST(wr.Path, wr.Handler)
	case "PUT":
		g.PUT(wr.Path, wr.Handler)
	case "DELETE":
		g.DELETE(wr.Path, wr.Handler)
	}
}
