package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/google/uuid"
)

type AuthAccountRepository struct {
	db DBTX
}

// FindLocalByEmail implements [authaccount.Repository].
func (a *AuthAccountRepository) FindLocalByEmail(
	ctx context.Context,
	email string,
) (*authaccount.AuthAccount, error) {
	query := `
		SELECT
			id,
			user_id,
			provider,
			provider_user_id,
			email,
			email_verified,
			password_hash,
			is_primary,
			created_at,
			updated_at
		FROM auth_accounts
		WHERE email = $1
		  AND provider = $2
	`

	var account authaccount.AuthAccount

	err := a.db.QueryRowContext(
		ctx,
		query,
		email,
		authaccount.ProviderLocal,
	).Scan(
		&account.ID,
		&account.UserID,
		&account.Provider,
		&account.ProviderUserID,
		&account.Email,
		&account.EmailVerified,
		&account.PasswordHash,
		&account.IsPrimary,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, authaccount.ErrAuthAccountNotFound
	}

	if err != nil {
		fmt.Println("Repo -> FindLocalByEmail -> ERROR: ", err)
		return nil, err
	}

	return &account, nil
}

// Create
func (a AuthAccountRepository) Create(ctx context.Context, account *authaccount.CreateInput) (*uuid.UUID, error) {
	query := `
		INSERT INTO auth_accounts (
			user_id, provider, provider_user_id, email, password_hash,is_primary
		) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id
	`
	var id uuid.UUID
	err := a.db.QueryRowContext(ctx, query,
		account.UserID,
		account.Provider,
		account.ProviderUserID,
		account.Email,
		account.PasswordHash,
		account.IsPrimary,
	).Scan(&id)

	return &id, err
}

// Delete implements [authaccount.Repository].
func (a *AuthAccountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// FindByID implements [authaccount.Repository].
func (a *AuthAccountRepository) FindByID(ctx context.Context, id uuid.UUID) (*authaccount.AuthAccount, error) {
	query := `
		SELECT id, user_id, provider, provider_user_id, email,email_verified, password_hash,is_primary, created_at, updated_at,last_login_at,
		linked_at
		FROM auth_accounts
		WHERE id = $1
	`

	var c authaccount.AuthAccount

	err := a.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.UserID,
		&c.Provider,
		&c.ProviderUserID,
		&c.Email,
		&c.EmailVerified,
		&c.PasswordHash,
		&c.IsPrimary,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.LastLoginAt,
		&c.LinkedAt,
	)

	if err == sql.ErrNoRows {
		return nil, authaccount.ErrAuthAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// FindByProvider implements [authaccount.Repository].
func (a *AuthAccountRepository) FindByProvider(ctx context.Context, provider authaccount.Provider, providerUserID string) (*authaccount.AuthAccount, error) {
	query := `
		SELECT id, user_id, provider, provider_user_id, email,email_verified, password_hash,is_primary, created_at, updated_at,last_login_at,
		linked_at
		FROM auth_accounts
		WHERE provider_user_id = $1 AND provider = $2
	`

	var c authaccount.AuthAccount

	err := a.db.QueryRowContext(ctx, query, providerUserID, provider).Scan(
		&c.ID,
		&c.UserID,
		&c.Provider,
		&c.ProviderUserID,
		&c.Email,
		&c.EmailVerified,
		&c.PasswordHash,
		&c.IsPrimary,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.LastLoginAt,
		&c.LinkedAt,
	)

	if err == sql.ErrNoRows {
		return nil, authaccount.ErrAuthAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// FindLocalByEmail implements [authaccount.Repository].
func (a *AuthAccountRepository) FindByEmail(ctx context.Context, email string) ([]*authaccount.AuthAccount, error) {
	var accounts []*authaccount.AuthAccount

	query := `
		SELECT id, user_id, provider, provider_user_id, email,email_verified, password_hash,is_primary, created_at, updated_at,last_login_at,
		linked_at
		FROM auth_accounts
		WHERE email = $1
	`

	rows, err := a.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account authaccount.AuthAccount

		err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Provider,
			&account.ProviderUserID,
			&account.Email,
			&account.EmailVerified,
			&account.PasswordHash,
			&account.IsPrimary,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.LastLoginAt,
			&account.LinkedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if len(accounts) == 0 {
		return nil, authaccount.ErrAuthAccountNotFound
	}

	return accounts, nil
}

// ListByUserID implements [authaccount.Repository].
func (a *AuthAccountRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*authaccount.AuthAccount, error) {
	var accounts []*authaccount.AuthAccount

	query := `
		SELECT id, user_id, provider, provider_user_id, email,email_verified, password_hash,is_primary, created_at, updated_at,last_login_at,
		linked_at
		FROM auth_accounts
		WHERE user_id = $1
	`

	rows, err := a.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account authaccount.AuthAccount

		err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Provider,
			&account.ProviderUserID,
			&account.Email,
			&account.EmailVerified,
			&account.PasswordHash,
			&account.IsPrimary,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.LastLoginAt,
			&account.LinkedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if len(accounts) == 0 {
		return nil, authaccount.ErrAuthAccountNotFound
	}

	return accounts, nil
}

// Update implements [authaccount.Repository].
func (a *AuthAccountRepository) Update(ctx context.Context, account *authaccount.AuthAccount) error {
	query := `
		UPDATE auth_accounts
		SET
			email = $1,
			password_hash = $2,
			email_verified = $3,
			is_primary = $4,
			last_login_at = $5,
			updated_at = $6
		WHERE id = $7
	`
	_, err := a.db.ExecContext(ctx, query,
		account.Email,
		account.PasswordHash,
		account.EmailVerified,
		account.IsPrimary,
		account.LastLoginAt,
		account.UpdatedAt,
		account.ID,
	)

	if err != nil {
		fmt.Printf("UPDATE USER ERR: %s", err.Error())
		return err
	}

	return nil
}

func NewAuthAccountRepository(db DBTX) authaccount.Repository {
	return &AuthAccountRepository{db: db}
}
