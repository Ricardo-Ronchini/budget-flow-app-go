package service

import (
	"fmt"
	"time"

	"github.com/ricardo-ronchini/budget-flow-app-go/common"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       string    `json:"user_id"`
	Name         *string   `json:"name"`
	Email        *string   `json:"email"`
	UserName     *string   `json:"username"`
	PasswordHash *string   `json:"password_hash"`
	Password     *string   `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginUser struct {
	UserName *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
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

func (data *LoginUser) LoginValidation() error {
	if data == nil {
		return fmt.Errorf("empty data")
	}

	if data.UserName != nil && *data.UserName == "" {
		return fmt.Errorf("invalid username")
	}

	return nil
}

func (data *LoginUser) GetUserForLogin(ctx *contexts.Context, db DB) (*User, error) {
	ctx.Logs().Logger.Debug("[GET USER LOGIN] find user to login: ", *data.UserName, ":", *data.Password)

	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	if err := data.LoginValidation(); err != nil {
		return nil, err
	}

	args := []any{
		data.UserName,
	}

	row := db.QueryRow(GetUserForLoginQuery, args...)

	var user User

	if err := row.Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.UserName,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		ctx.Logs().Logger.Debug("error parsing user: ", err)
		return nil, err
	}

	return &user, nil
}

func (data *User) CreateUser(ctx *contexts.Context, db DB) error {
	if db == nil {
		return fmt.Errorf("connection not established")
	}

	ctx.Logs().Logger.Debug("check validations to create user")

	if err := data.Validation(); err != nil {
		return err
	}

	if data.Password == nil {
		return fmt.Errorf("missing password content")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generating hash password")
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
		return fmt.Errorf("error creating user")
	}

	return nil
}

func (data *User) GetUserByID(ctx *contexts.Context, db DB) (*User, error) {
	if db == nil {
		return nil, fmt.Errorf("connection not established")
	}

	if data.UserID == "" {
		return nil, fmt.Errorf("user id cannot be empty")
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
