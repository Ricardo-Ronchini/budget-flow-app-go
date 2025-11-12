package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ricardo-ronchini/budget-flow-app-go/common"
)

type Expense struct {
	ExpenseID string    `json:"expense_id"`
	Name      string    `json:"name"`
	Value     float32   `json:"value"`
	Date      time.Time `json:"date"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (data *Expense) Validation() error {
	if data.ExpenseID == "" {
		return fmt.Errorf("expense id cannot be empty")
	}

	return nil
}

func (data *Expense) CreateValidation() error {
	if data.Name == "" {
		return fmt.Errorf("expense name cannot be empty")
	}

	if data.UserID == "" {
		return fmt.Errorf("user id cannot be empty")
	}

	return nil
}

func Expenses(db DB) (*[]Expense, error) {
	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	rows, err := db.Query(GetAllExpensesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense

	for rows.Next() {
		var expense Expense

		if err := rows.Scan(
			&expense.ExpenseID,
			&expense.Name,
			&expense.Value,
			&expense.Date,
			&expense.UserID,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		); err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return &expenses, nil
}

func (data *Expense) ExpenseByID(db DB) (*Expense, error) {
	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	if err := data.Validation(); err != nil {
		return nil, err
	}

	rows := db.QueryRow(GetExpenseByIDQuery, data.ExpenseID)

	var expense Expense

	if err := rows.Scan(
		&expense.ExpenseID,
		&expense.Name,
		&expense.Value,
		&expense.Date,
		&expense.UserID,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &expense, nil
}

func (data *Expense) CreateExpense(db DB) error {
	if db == nil {
		return fmt.Errorf("connection not established")
	}

	if err := data.CreateValidation(); err != nil {
		return err
	}

	expenseID := common.GenerateCustomGuideID()

	args := []any{
		expenseID,
		data.Name,
		data.Value,
		data.Date,
		data.UserID,
		time.Now(),
		time.Now(),
	}

	_, err := db.Exec(CreateExpenseQuery, args...)

	return err
}

func (data *Expense) UpdateExpense(db DB) error {
	if db == nil {
		return fmt.Errorf("connection not established")
	}

	if err := data.Validation(); err != nil {
		return err
	}

	expense, err := data.ExpenseByID(db)
	if err != nil {
		return fmt.Errorf("error when get expense. Error: %v", err)
	}

	if expense == nil {
		return fmt.Errorf("expense not found")
	}

	args := []any{
		data.ExpenseID,
		data.Name,
		data.Value,
		data.Date,
		data.UserID,
		time.Now(),
	}

	_, err = db.Exec(UpdateExpenseQuery, args...)

	return err
}

func DeleteExpense(expenseID string, db DB) error {
	if db == nil {
		return fmt.Errorf("connection not established")
	}

	if expenseID == "" {
		return fmt.Errorf("expense id cannot be empty")
	}

	data := Expense{ExpenseID: expenseID}

	expense, err := data.ExpenseByID(db)
	if err != nil {
		return fmt.Errorf("error when get expense. Error: %v", err)
	}

	if expense == nil {
		return fmt.Errorf("expense not found")
	}

	_, err = db.Exec(DeleteExpenseQuery, expenseID)

	return err
}
