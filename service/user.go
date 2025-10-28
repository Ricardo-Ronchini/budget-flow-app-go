package service

import (
	"fmt"
	"time"

	"github.com/ricardo-ronchini/budget-flow-app-go/common"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    string    `json:"user_id"`
	Name      *string   `json:"name"`
	Email     *string   `json:"email"`
	UserName  *string   `json:"user_name"`
	Password  *string   `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	nameMaxChar int = 30
	nameMinChar int = 4
)

func (data *User) Validation() error {
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

func (data *User) GetUserForLogin(ctx *contexts.Context, db DB) (*User, error) {
	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	if err := data.Validation(); err != nil {
		return nil, err
	}

	args := []any{
		data.UserName,
	}

	rows, err := db.Query(GetUserForLoginQuery, args)
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
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (data *User) CreateUser(ctx *contexts.Context, db DB) error {
	if db == nil {
		return fmt.Errorf("connection not established")
	}

	if err := data.Validation(); err != nil {
		return err
	}

	if data.Password == nil {
		return fmt.Errorf("senha inválida")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar senha hash")
	}

	passwordHash := string(bytes)

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

	if _, err := db.Exec(CreateUserQuery, args...); err != nil {
		return fmt.Errorf("erro ao inserir usuário")
	}

	return nil
}

func (data *User) GetUserByID(ctx *contexts.Context, db DB) (*User, error) {
	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	args := []any{
		data.UserID,
	}

	rows := db.QueryRow(GetUserByIDQuery, args...)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var user User

	if err := rows.Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.UserName,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
