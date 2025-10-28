package service

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ricardo-ronchini/budget-flow-app-go/common"
)

const (
	mockExpenseID string = "1"
	mockUserID    string = "2"
	mockName      string = "Gasto 1"
)

var (
	mockValue     float32   = 19.90
	mockDate      time.Time = time.Date(2024, 9, 1, 11, 0, 0, 0, time.UTC)
	mockCreatedAt time.Time = time.Date(2024, 10, 2, 15, 30, 0, 0, time.UTC)
	mockUpdatedAt time.Time = time.Date(2024, 10, 3, 15, 30, 0, 0, time.UTC)
)

func TestExpenses(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"expense_id", "name", "value", "date", "user_id", "created_at", "updated_at",
	}).AddRow(mockExpenseID, mockName, mockValue, mockDate, mockUserID, mockCreatedAt, mockUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(GetAllExpensesQuery)).
		WillReturnRows(rows)

	result, err := Expenses(db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
	}

	if result == nil {
		t.Error("expected result not nil")
		return
	}

	if len(*result) < 1 {
		t.Error("expected result greater than zero")
	}

	for _, expense := range *result {
		if expense.Name != mockName {
			t.Errorf("expected name '%s', got '%s'", mockName, expense.Name)
		}
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGetExpenseByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"expense_id", "name", "value", "date", "user_id", "created_at", "updated_at",
	}).AddRow(mockExpenseID, mockName, mockValue, mockDate, mockUserID, mockCreatedAt, mockUpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(GetExpenseByIDQuery)).
		WithArgs(mockExpenseID).
		WillReturnRows(rows)

	expense := Expense{ExpenseID: mockExpenseID}

	result, err := expense.ExpenseByID(db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
		return
	}

	if result == nil {
		t.Error("expected result not nil")
		return
	}

	if result.Name != mockName {
		t.Errorf("expected name '%s', got '%s'", mockName, expense.Name)
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestCreateExpense(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(CreateExpenseQuery)).
		WithArgs(mockExpenseID, mockName, mockValue, mockDate, mockUserID, mockCreatedAt, mockUpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	expense := Expense{
		ExpenseID: mockExpenseID,
		Name:      mockName,
		Value:     mockValue,
		Date:      mockDate,
		UserID:    mockUserID,
		CreatedAt: mockCreatedAt,
		UpdatedAt: mockUpdatedAt,
	}

	err := expense.CreateExpense(db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
		return
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestUpdateExpense(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(UpdateExpenseQuery)).
		WithArgs(mockExpenseID, mockName, mockValue, mockDate, mockUserID, mockUpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	expense := Expense{
		ExpenseID: mockExpenseID,
		Name:      mockName,
		Value:     common.RoundTo(mockValue, 2),
		Date:      mockDate,
		UserID:    mockUserID,
		UpdatedAt: mockUpdatedAt,
	}

	err := expense.UpdateExpense(db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
		return
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestDeleteExpense(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(DeleteExpenseQuery)).
		WithArgs(mockExpenseID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteExpense(mockExpenseID, db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
		return
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
