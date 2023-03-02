package domain

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type User struct {
	ID             string
	Username       string
	HashedPassword string
	Email          *string
	IsActive       bool
	CreatedAt      time.Time
	ModifiedAt     *time.Time
	LastLoginAt    *time.Time
	SessionID      string
}

type UserPages struct {
	UserID string
	Pages  []UserPages
}

type UserPage struct {
	URL       string
	CreatedAt time.Time
}
