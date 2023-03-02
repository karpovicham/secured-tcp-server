package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/karpovicham/secured-tcp-server/service/server/domain"
)

type repo struct {
	db    *pgx.Conn
	nowFn func() time.Time
}

func NewRepo(db *pgx.Conn) *repo {
	return &repo{
		db: db,
		nowFn: func() time.Time {
			return time.Now().UTC()
		},
	}
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	const query = `
		SELECT id, username, hashed_password
		FROM users
		WHERE is_active = true AND username = $1
	`

	var user domain.User
	if err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.HashedPassword,
	); err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, domain.ErrNotFound
		}

		return domain.User{}, err
	}

	return user, nil
}

func (r *repo) GetUserByUserID(ctx context.Context, userID string) (domain.User, error) {
	const query = `
		SELECT id, username, email, last_login_at, session_id
		FROM users
		WHERE is_active = true AND id = $1
	`

	var user domain.User
	if err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.LastLoginAt,
		&user.SessionID,
	); err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, domain.ErrNotFound
		}

		return domain.User{}, err
	}

	return user, nil
}

func (r *repo) UpdateUserSession(ctx context.Context, username, sessionID string, lastLoginAt time.Time) error {
	const query = `
		UPDATE users 
		SET session_id = $1, last_login_at = $2
		WHERE username = $3
	`

	_, err := r.db.Exec(ctx, query, sessionID, lastLoginAt, username)
	if err != nil {
		return err
	}

	return nil
}

// ClearUserSession make  session ID empty for the user
func (r *repo) ClearUserSession(ctx context.Context, userID string) error {
	const query = `
		UPDATE users 
		SET session_id = NULL
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *repo) DeactivateUser(ctx context.Context, userID string) error {
	const query = `
		UPDATE users 
		SET is_active = false, modified_at = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(ctx, query, r.nowFn(), userID)
	return err
}

func (r *repo) UpdateUserData(ctx context.Context, userID, username string, email *string) error {
	const query = `
		UPDATE users 
		SET username = $1, email = $2, modified_at = $3
		WHERE id = $4
	`

	// TODO: check on unique error
	_, err := r.db.Exec(ctx, query, username, email, r.nowFn(), userID)
	return err
}

// AddNewUserPage Add new page and indicate that account setting is being updated
func (r *repo) AddNewUserPage(ctx context.Context, userID, pageURL string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	const pageQuery = `
		INSERT INTO user_pages (user_id, url) 
		VALUES ($1, $2)
	`

	if _, err = tx.Exec(ctx, pageQuery, userID, pageURL); err != nil {
		return fmt.Errorf("exec page query: %w", err)
	}

	const userQuery = `
		UPDATE users 
		SET modified_at = $1
		WHERE id = $4
	`

	if _, err = tx.Exec(ctx, userQuery, userID, r.nowFn()); err != nil {
		return fmt.Errorf("exec user query: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
