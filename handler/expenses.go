package handler

import (
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/service"
)

const exepensePath string = "/expenses"

var V1ExpensesGET = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.GET,
	Handler: func(c *contexts.Context) (int, any) {
		db := c.Database().Connect()
		defer db.Close()

		result, err := service.Expenses(db)
		if err != nil {
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		return http.StatusOK, result
	},
}

var V1EspensesByIDGET = &contexts.WebRoute{
	Path:   exepensePath + "/:expense_id",
	Method: contexts.GET,
	Handler: func(c *contexts.Context) (int, any) {
		expenseID := c.EchoContext.QueryParam("expense_id")

		data := service.Expense{
			ExpenseID: expenseID,
		}

		db := c.Database().Connect()
		defer db.Close()

		result, err := data.ExpenseByID(db)
		if err != nil {
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		return http.StatusOK, result
	},
}

var V1ExpensesPOST = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {
		data := service.Expense{}

		c.EchoContext.Bind(&data)

		tx, err := c.Database().Connect().Begin()
		if err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		if err = data.CreateExpense(tx); err != nil {
			tx.Rollback()
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		if err = tx.Commit(); err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		return http.StatusOK, nil
	},
}

var V1ExpensesPUT = &contexts.WebRoute{
	Path:   exepensePath + "/:expense_id",
	Method: contexts.PUT,
	Handler: func(c *contexts.Context) (int, any) {
		expenseID := c.EchoContext.QueryParam("expense_id")
		data := service.Expense{ExpenseID: expenseID}

		c.EchoContext.Bind(&data)

		tx, err := c.Database().Connect().Begin()
		if err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		if err := data.UpdateExpense(tx); err != nil {
			tx.Rollback()
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		if err = tx.Commit(); err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		return http.StatusOK, nil
	},
}

var V1ExpensesDELETE = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.DELETE,
	Handler: func(c *contexts.Context) (int, any) {
		expenseID := c.EchoContext.QueryParam("expense_id")

		tx, err := c.Database().Connect().Begin()
		if err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())

		}

		if err := service.DeleteExpense(expenseID, tx); err != nil {
			tx.Rollback()
			return http.StatusBadRequest, c.API().Error(http.StatusBadRequest, err.Error())
		}

		if err = tx.Commit(); err != nil {
			return http.StatusInternalServerError, c.API().Error(http.StatusInternalServerError, err.Error())
		}

		return http.StatusOK, nil
	},
}
