package service

import (
	"context"
	"fmt"
	"time"
)

type Expense struct {
	Name      string    `json:"name"`
	Value     float32   `json:"value"`
	Date      time.Time `json:"date"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Services) Expenses(ctx context.Context) (*[]Expense, error) {
	query := `SELECT name, value, date, user_id, created_at FROM expense`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense

	for rows.Next() {
		var expense Expense

		if err := rows.Scan(&expense.Name, &expense.Value, &expense.Date, &expense.UserID, &expense.CreatedAt); err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return &expenses, nil
}

func (s *Services) ExpenseByID(ctx context.Context, expenseID string) (*Expense, error) {
	if expenseID == "" {
		return nil, fmt.Errorf("expense id cannot be empty")
	}

	query := fmt.Sprintf(`SELECT name, value, date, user_id, created_at FROM expense WHERE expense_id = %s`, expenseID)

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expense Expense

	if err := rows.Scan(&expense.Name, &expense.Value, &expense.Date, &expense.UserID, &expense.CreatedAt); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (s *Services) CreateExpense(ctx context.Context, expenseData Expense) error {
	query := `INSERT INTO expense (name, value, date, user_id, created_at) VALUES ($1, $2, $3, $4, $5)`

	args := []any{
		expenseData.Name,
		expenseData.Value,
		expenseData.Date,
		expenseData.UserID,
		time.Now(),
	}

	_, err := s.DB.ExecContext(ctx, query, args...)

	return err
}
