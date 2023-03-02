package server

import (
	"context"
	"time"

	"github.com/karpovicham/secured-tcp-server/service/server/domain"
)

type Repo interface {
	// GetUserByUsername returns user with ID and Hashed Password
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)

	// GetUserByUserID returns user data by its ID
	GetUserByUserID(ctx context.Context, userID string) (domain.User, error)

	// UpdateUserSession saves session data for the user
	UpdateUserSession(ctx context.Context, username, sessionID string, lastLoginAt time.Time) error

	// ClearUserSession clear session data for the user
	ClearUserSession(ctx context.Context, userID string) error

	// DeactivateUser switch user active flag to false
	DeactivateUser(ctx context.Context, userID string) error

	// UpdateUserData update user email/username
	UpdateUserData(ctx context.Context, userID, username string, email *string) error

	// AddNewUserPage add new user page
	AddNewUserPage(ctx context.Context, userID, pageURL string) error
}
