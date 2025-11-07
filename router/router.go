package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ricardo-ronchini/budget-flow-app-go/auth"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/handler"
)

type Routes *contexts.EchoHandler

func Init(e *echo.Echo) {
	// *** Echo ***
	c := contexts.NewContext()

	c.Logs().Info("[ROUTER] Init API services")

	e.Logger.SetLevel(1)
	e.Logger.SetOutput(os.Stdout)

	e.Use(auth.ConfigCORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | [${method}] '${uri}' -> ${status} (${latency_human})\n",
	}))

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "API service is working! 🚀")
	})

	v1 := e.Group("/v1")
	api := v1.Group("/api", auth.Middleware)

	noAuthRoutes := []Routes{
		c.HandlerWebRoute(handler.V1Login),
	}

	authRoutes := []Routes{
		// user
		c.HandlerWebRoute(handler.V1UserGET),
		c.HandlerWebRoute(handler.V1UserPOST),
		// expense
		c.HandlerWebRoute(handler.V1ExpensesGET),
		c.HandlerWebRoute(handler.V1EspensesByIDGET),
		c.HandlerWebRoute(handler.V1ExpensesPOST),
		c.HandlerWebRoute(handler.V1ExpensesPUT),
		c.HandlerWebRoute(handler.V1ExpensesDELETE),
	}

	c.Logs().Info("[ROUTER] Recording routes...")

	for _, r := range noAuthRoutes {
		register(v1, r)
	}

	for _, r := range authRoutes {
		register(api, r)
	}

	for _, r := range e.Routes() {
		c.Logs().Info(r.Method, " -> ", r.Path)
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
