package contexts

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/common"
	"github.com/ricardo-ronchini/budget-flow-app-go/db"
)

type Context struct {
	Database func() *db.Database
	API      func() *APIClient
	Logs     func() *Logs

	Host          string
	APIVersion    string
	EchoContext   echo.Context
	SystemContext context.Context
}

func NewContext() *Context {
	return &Context{
		Host:       common.GetEnv("HOST", "localhost:8080"),
		APIVersion: common.GetEnv("API_VERSION", "V.0"),
		Logs:       NewLogs,
		API:        func() *APIClient { return nil },
		Database:   func() *db.Database { return nil },
	}
}
