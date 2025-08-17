package handler

import (
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/service"
)

const userPath string = "/user"

var V1UserGET = &contexts.WebRoute{
	Path:   userPath + "/:user_id",
	Method: contexts.GET,
	Handler: func(c *contexts.Context) (int, any) {
		db := c.Database().Connect()
		defer db.Close()

		_ = db

		user := []string{"Ricardo", "Amanda"}

		return http.StatusOK, c.API().Ok(user)
	},
}

var V1UserPOST = &contexts.WebRoute{
	Path:   userPath,
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {
		db := c.Database().Connect()
		defer db.Close()

		// services := service.NewServiceContext(db)

		user := service.User{}

		return http.StatusOK, c.API().Ok(user)
	},
}
