package service

import "database/sql"

// DB is the minimal database interface accepted by all service functions.
// Using an interface instead of *sql.DB allows tests to substitute a mock without a real connection.
type DB interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}
