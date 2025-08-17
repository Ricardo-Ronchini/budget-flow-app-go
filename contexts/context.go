package contexts

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/db"
)

// informaçoes de sessão (user, token, nivel de acesso, idioma, fuso, etc)
// logs
// conexao banco de dados

type Context struct {
	Database func() *db.Database
	API      func() *APIClient
	Logs     func() *Logs

	Host        string
	APIVersion  string
	EchoContext echo.Context
}

func NewContext() *Context {
	return &Context{
		Host:       os.Getenv("HOST"),
		APIVersion: os.Getenv("API_VERSION"),
	}
}

func test() {
	// consumindo API's call
	c := NewContext()
	c.API().Ok(nil)
	c.API().Error(0, "")

	c.API().Session()

	c.Database().Connect() // Ok
	db := c.Database().Connect()
	_, _ = db.Begin()
}
