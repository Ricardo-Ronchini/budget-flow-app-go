package contexts

import "github.com/labstack/echo/v4"

type httpMethod string

const (
	GET    httpMethod = "GET"
	PATCH  httpMethod = "PATCH"
	PUT    httpMethod = "PUT"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
)

type EchoHandler struct {
	Path       string
	Method     httpMethod
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type WebRoute struct {
	Path       string
	Method     httpMethod
	Handler    func(c *Context) (int, any)
	Middleware []echo.MiddlewareFunc
}

func (ctx *Context) HandlerWebRoute(w *WebRoute) *EchoHandler {
	return &EchoHandler{
		Path:   w.Path,
		Method: w.Method,
		Handler: func(ec echo.Context) error {
			ctx.EchoContext = ec
			code, resp := w.Handler(ctx)
			return ec.JSON(code, resp)
		},
		Middleware: w.Middleware,
	}
}
