package contexts

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

type httpMethod string

const (
	GET    httpMethod = "GET"
	PATCH  httpMethod = "PATCH"
	PUT    httpMethod = "PUT"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
)

type PermissionName string

const (
	AdminLevel  PermissionName = "admin"
	MediumLevel PermissionName = "medium"
	BasicLevel  PermissionName = "basic"
)

type EchoHandler struct {
	Path       string
	Method     httpMethod
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type WebRoute struct {
	Path            string
	Method          httpMethod
	Handler         func(c *Context) (int, any)
	Middleware      []echo.MiddlewareFunc
	PermissionLevel []PermissionName
}

func (ctx *Context) CheckPermissionLevel(w *WebRoute) bool {
	return slices.Contains(w.PermissionLevel, PermissionName(ctx.API().Session().UserLevel))
}

func (ctx *Context) HandlerWebRoute(w *WebRoute) *EchoHandler {
	return &EchoHandler{
		Path:   w.Path,
		Method: w.Method,
		Handler: func(ec echo.Context) error {
			ctx.EchoContext = ec

			if allowed := ctx.CheckPermissionLevel(w); !allowed {
				return ec.JSON(http.StatusForbidden, map[string]string{
					"error": "permissão insuficiente",
				})
			}

			code, resp := w.Handler(ctx)
			return ec.JSON(code, resp)
		},
		Middleware: w.Middleware,
	}
}
