package service

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"user_id", "name", "email", "user_name", "created_at", "modified_at",
	}).AddRow("1", "Ricardo", "ricardo@email.com", "ricardodev", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(GetUserByIDQuery)).
		WithArgs("1").
		WillReturnRows(rows)

	user := &User{UserID: "1"}

	result, err := user.GetUserByID(nil, db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
	}

	if result == nil {
		t.Error("expected result not nil...")
	}

	if result.Name == nil {
		t.Error("expected name value is != nil...")
	}

	if *result.Name != "Ricardo" {
		t.Errorf("expected name Ricardo, got %s", *result.Name)
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"user_id", "name", "email", "user_name", "created_at", "modified_at",
	}).AddRow("1", "Ricardo", "ricardo@email.com", "ricardodev", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(CreateUserQuery)).
		WithArgs("1").
		WillReturnRows(rows)

	user := &User{UserID: "1"}

	err := user.CreateUser(nil, db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestGetUserForLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"user_id", "name", "email", "user_name", "created_at", "modified_at",
	}).AddRow("1", "Ricardo", "ricardo@email.com", "ricardodev", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(GetUserForLoginQuery)).
		WithArgs("1").
		WillReturnRows(rows)

	user := &User{UserID: "1"}

	result, err := user.GetUserByID(nil, db)

	if err != nil {
		t.Errorf("query exec error: %v", err)
	}

	if result == nil {
		t.Error("expected result not nil...")
	}

	if result.Name == nil {
		t.Error("expected name value is != nil...")
	}

	if *result.Name != "Ricardo" {
		t.Errorf("expected name Ricardo, got %s", *result.Name)
	}

	// validates if all tests passed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
