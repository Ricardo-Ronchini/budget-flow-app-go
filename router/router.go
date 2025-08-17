package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/auth"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/handler"
)

type Routes *contexts.EchoHandler

func Init(e *echo.Echo) {
	// *** Echo ***
	c := contexts.NewContext()

	// no auth
	e.POST("login", handler.Login)

	// add middleware
	api := e.Group("/api", auth.Middleware)

	// middleware
	// - CORS
	// - Auth
	// - Level

	// auth's
	routes := []Routes{
		c.HandlerWebRoute(handler.V1UserGET),
		// user
		// handler.V1UserGET,
		// // expense
		// handler.V1ExpensesGET,
		// handler.V1ExpensesPOST,
		// handler.V1ExpensesPUT,
		// handler.V1ExpensesPATCH,
		// handler.V1ExpensesDELETE,
	}

	for _, r := range routes {
		register(api, r)
	}
}

func register(g *echo.Group, wr Routes) {
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
