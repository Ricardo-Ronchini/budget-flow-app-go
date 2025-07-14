package service

import "database/sql"

type Services struct {
	DB *sql.DB
}

func NewServiceContext(db *sql.DB) *Services {
	return &Services{
		DB: db,
	}
}
