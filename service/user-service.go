package service

import (
	"context"
	"time"

	"github.com/ricardo-ronchini/budget-flow-app-go/common"
)

type User struct {
	UserID     string    `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	UserName   string    `json:"user_name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (s *Services) UserTokenByUserName(ctx context.Context) (*User, error) {
	query := `
		SELECT 
			user_id, name, email, user_name, created_at, modified_at
		FROM 
			users
		WHERE
			user_name = $1
	`

	args := []any{
		"",
	}

	rows, err := s.DB.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Err() != nil {
		return nil, err
	}

	var user User

	if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.UserName, &user.CreatedAt, &user.ModifiedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Services) CreateUser(ctx context.Context, dataUser *User) error {
	// validations

	query := `
		INSERT into users 
			(user_id, name, email, user_name, created_at, modified_at)
		VALUES 
			($1, $2, $3, $4, $5, $6)
		;
	`

	userID := common.GenerateCustomGuideID()

	args := []any{
		userID,
		dataUser.Name,
		dataUser.Email,
		dataUser.UserName,
		time.Now(),
		time.Now(),
	}

	_, err := s.DB.ExecContext(ctx, query, args...)

	return err
}
