package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/config"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
)

const exepensePath string = "/expenses"

var V1ExpensesGET = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.GET,
	Handler: func(c echo.Context) error {
		db, err := config.ConnectDB()
		if err != nil {
			log.Fatal("Erro ao conectar no banco:", err)
		}
		defer db.Close()

		// Chama o serviço (simulado)
		expenses := []string{"Aluguel", "Comida"}

		// Retorna JSON
		return c.JSON(200, echo.Map{
			"expenses": expenses,
		})
	},
}

var V1EspensesByIDGET = &contexts.WebRoute{
	Path:   exepensePath + "/:expense_id",
	Method: contexts.GET,
	Handler: func(c echo.Context) error {
		expenseID := c.QueryParam("expense_id")
		expenses := []string{"Aluguel", "Comida"}

		_ = expenseID
		return c.JSON(200, echo.Map{
			"data": expenses,
		})
	},
}

var V1ExpensesPOST = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.POST,
	Handler: func(c echo.Context) error {

		return c.JSON(200, echo.Map{
			"data": "create new expense",
		})
	},
}

var V1ExpensesPUT = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.PUT,
	Handler: func(c echo.Context) error {

		return c.JSON(200, echo.Map{
			"data": "update expense",
		})
	},
}

var V1ExpensesPATCH = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.PATCH,
	Handler: func(c echo.Context) error {

		return c.JSON(200, echo.Map{
			"data": "change exchange",
		})
	},
}

var V1ExpensesDELETE = &contexts.WebRoute{
	Path:   exepensePath,
	Method: contexts.DELETE,
	Handler: func(c echo.Context) error {

		return c.JSON(200, echo.Map{
			"data": "delete exchange",
		})
	},
}
