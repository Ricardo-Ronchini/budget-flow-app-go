package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // driver PostgreSQL
)

type DBConnection struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DB_NAME  string
}

func ContextDB() DBConnection {
	return DBConnection{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:  os.Getenv("DB_NAME"),
	}
}

func ConnectDB() (*sql.DB, error) {
	dbContext := ContextDB()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbContext.HOST, dbContext.PORT, dbContext.USER, dbContext.PASSWORD, dbContext.DB_NAME,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// ping connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("ping", db.Ping())

	return db, nil
}
