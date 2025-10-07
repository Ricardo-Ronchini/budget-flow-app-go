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
		userID := c.EchoContext.QueryParam("user_id")
		user := service.User{UserID: userID}

		response, err := user.GetUserByID(c, nil)
		if err != nil {
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		return http.StatusOK, c.API().Ok(response)
	},
	PermissionLevel: []contexts.PermissionName{contexts.AdminLevel, contexts.MediumLevel, contexts.BasicLevel},
}

var V1UserPOST = &contexts.WebRoute{
	Path:   userPath,
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {
		user := service.User{}
		c.EchoContext.Bind(&user)

		db := c.Database().Connect()
		defer db.Close()

		tx, err := db.Begin()
		if err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		if err := user.CreateUser(c, tx); err != nil {
			tx.Rollback()
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		if err = tx.Commit(); err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		return http.StatusOK, c.API().Ok(user)
	},
	PermissionLevel: []contexts.PermissionName{contexts.AdminLevel, contexts.MediumLevel},
}
