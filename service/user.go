package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ricardo-ronchini/budget-flow-app-go/common"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID     string    `json:"user_id"`
	Name       *string   `json:"name"`
	Email      *string   `json:"email"`
	UserName   *string   `json:"user_name"`
	Password   *string   `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

const (
	nameMaxChar int = 30
	nameMinChar int = 4
)

func Validation(data *User) error {
	if data == nil {
		return fmt.Errorf("empty data")
	}

	if data.UserName != nil && *data.UserName == "" {
		return fmt.Errorf("invalid username")
	}

	if data.Email != nil && *data.Email == "" {
		return fmt.Errorf("invalid email")
	}

	if data.Name != nil {
		if *data.Name == "" {
			return fmt.Errorf("invalid name")
		}

		if len(*data.Name) > nameMaxChar {
			return fmt.Errorf("name exceeded the 30 character limit")
		}

		if len(*data.Name) < nameMinChar {
			return fmt.Errorf("name must contain more than 4 characters")
		}
	}

	return nil
}

func (data *User) GetUserForLogin(ctx *contexts.Context, tx *sql.DB) (*User, error) {
	if err := Validation(data); err != nil {
		return nil, err
	}

	if tx == nil {
		tx = ctx.Database().Connect()
	}
	defer tx.Close()

	query := `
		SELECT 
			user_id, name, email, user_name, created_at, modified_at
		FROM 
			users
		WHERE
			user_name = $1
		OR
			email = $1
		;
	`

	args := []any{
		data.UserName,
	}

	rows, err := tx.Query(query, args)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	var user User

	if err := rows.Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.UserName,
		&user.CreatedAt,
		&user.ModifiedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (data *User) CreateUser(ctx *contexts.Context, tx *sql.Tx) error {
	if err := Validation(data); err != nil {
		return err
	}

	if tx == nil {
		return fmt.Errorf("transação nao definida")
	}

	if data.Password == nil {
		return fmt.Errorf("senha inválida")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar senha hash")
	}

	passwordHash := string(bytes)

	query := `
		INSERT into users 
			(user_id, name, email, user_name, password, created_at, modified_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7)
		;
	`

	userID := common.GenerateCustomGuideID()

	args := []any{
		userID,
		data.Name,
		data.Email,
		data.UserName,
		passwordHash,
		time.Now(),
		time.Now(),
	}

	if _, err := tx.Exec(query, args...); err != nil {
		return fmt.Errorf("erro ao inserir usuário")
	}

	return nil
}

func (data *User) GetUserByID(ctx *contexts.Context, tx *sql.DB) (*User, error) {
	if tx == nil {
		tx = ctx.Database().Connect()
	}
	defer tx.Close()

	query := `
		SELECT 
			user_id, name, email, user_name, created_at, modified_at
		FROM 
			users
		WHERE
			user_id = $1
		;
	`

	args := []any{
		data.UserName,
	}

	rows, err := tx.Query(query, args)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	var user User

	if err := rows.Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.UserName,
		&user.CreatedAt,
		&user.ModifiedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
