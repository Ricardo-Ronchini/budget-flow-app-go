package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Expense struct {
	ExpenseID string    `json:"expense_id"`
	Name      string    `json:"name"`
	Value     float32   `json:"value"`
	Date      time.Time `json:"date"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (data *Expense) Validation() error {
	if data.ExpenseID == "" {
		return fmt.Errorf("expense id cannot be empty")
	}

	return nil
}

func Expenses(db *sql.DB) (*[]Expense, error) {
	query := `SELECT name, value, date, user_id, created_at FROM expense`

	rows, err := db.Query(query)
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

func (data *Expense) ExpenseByID(db *sql.DB) (*Expense, error) {
	if err := data.Validation(); err != nil {
		return nil, err
	}

	query := `
		SELECT 
			name, value, date, user_id, created_at 
		FROM 
			expense 
		WHERE 
			expense_id = $1
		;
	`

	rows, err := db.Query(query, data.ExpenseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expense Expense

	if err := rows.Scan(
		&expense.Name,
		&expense.Value,
		&expense.Date,
		&expense.UserID,
		&expense.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (data *Expense) CreateExpense(tx *sql.Tx) error {
	query := `
		INSERT INTO expense 
			(expense_id, name, value, date, user_id, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6)
		;
	`

	args := []any{
		data.Name,
		data.Value,
		data.Date,
		data.UserID,
		time.Now(),
		time.Now(),
	}

	_, err := tx.Exec(query, args...)

	return err
}

func (data *Expense) UpdateExpense(tx *sql.Tx) error {
	if err := data.Validation(); err != nil {
		return err
	}

	query := `
		UPDATE expense
		SET
			name = $2,
			value = $3,
			updated_at = $4,
		WHERE
			expense_id = $1
		;
	`

	args := []any{
		data.Name,
		data.Value,
		time.Now(),
	}

	_, err := tx.Exec(query, args...)

	return err
}

func DeleteExpense(expenseID string, tx *sql.Tx) error {
	if expenseID == "" {
		return fmt.Errorf("expense id cannot be empty")
	}

	query := `
		DELETE expense
		WHERE
			expense_id = $1
		;
	`

	_, err := tx.Exec(query, expenseID)

	return err
}
