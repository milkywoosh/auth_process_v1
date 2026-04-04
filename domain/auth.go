package domain

import (
	"context"
	"time"
)

type User struct {
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
}

type Authentication interface {
	CreateUser(ctx context.Context, username string, password string) error
	GetUser(ctx context.Context, username string) (User, error)
}
