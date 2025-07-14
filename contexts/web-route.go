package contexts

import "github.com/labstack/echo/v4"

type WebRoute struct {
	Path       string
	Method     httpMethod
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type httpMethod string

const (
	GET    httpMethod = "GET"
	PATCH  httpMethod = "PATCH"
	PUT    httpMethod = "PUT"
	POST   httpMethod = "POST"
	DELETE httpMethod = "DELETE"
)
