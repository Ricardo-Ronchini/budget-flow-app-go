package handler

import (
	"net/http"

	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
)

const exepensePath string = "/expenses"

var V1ExpensesGET = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.GET,
	Handler: func(c *contexts.Context) (int, any) {
		db := c.Database().Connect()
		defer db.Close()

		// Chama o serviço (simulado)
		expenses := []string{"Aluguel", "Comida"}

		return http.StatusOK, expenses
	},
}

var V1EspensesByIDGET = &contexts.WebRoute{
	Path:   exepensePath + "/:expense_id",
	Method: contexts.GET,
	Handler: func(c *contexts.Context) (int, any) {
		expenseID := c.EchoContext.QueryParam("expense_id")
		expenses := []string{"Aluguel", "Comida"}

		_ = expenseID

		return http.StatusOK, expenses
	},
}

var V1ExpensesPOST = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {

		return http.StatusOK, nil
	},
}

var V1ExpensesPUT = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.PUT,
	Handler: func(c *contexts.Context) (int, any) {

		return http.StatusOK, nil
	},
}

var V1ExpensesPATCH = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.PATCH,
	Handler: func(c *contexts.Context) (int, any) {

		return http.StatusOK, nil
	},
}

var V1ExpensesDELETE = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.DELETE,
	Handler: func(c *contexts.Context) (int, any) {

		return http.StatusOK, nil
	},
}
