package repositories

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
)

type SessionRepository struct {
	db DBTX
}

// Desactive implements [session.Repository].
func (r *SessionRepository) Desactive(ctx context.Context, hash string) error {
	query := `
		UPDATE sessions SET
			actived = false,
			last_used_at = NOW()
		WHERE token_hash = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		hash,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return item.ErrItemNotFound
	}

	return nil
}

// Create implements [session.Repository].
func (r *SessionRepository) Create(ctx context.Context, session *session.Session) error {
	query := `
		INSERT INTO sessions
		(id, auth_account_id, token_hash, expires_at,created_at,last_used_at ) 
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.ExecContext(ctx, query,
		session.Id,
		session.AuthAccountId,
		session.TokenHash,
		session.ExpiresAt,
		session.CreatedAt,
		session.LastUsedAt,
	)

	return err
}

// Delete implements [session.Repository].
func (r *SessionRepository) Delete(ctx context.Context, hash string) error {
	query := `
		DELETE FROM sessions WHERE token_hash = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		hash,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return item.ErrItemNotFound
	}

	return nil
}

// FindById implements [session.Repository].
func (r *SessionRepository) FindByHash(ctx context.Context, hash string) (*session.Session, error) {
	query := `
		SELECT id, auth_account_id, token_hash, expires_at, actived
		FROM sessions
		WHERE token_hash = $1
	`

	var s session.Session

	err := r.db.QueryRowContext(ctx, query, hash).Scan(
		&s.Id,
		&s.AuthAccountId,
		&s.TokenHash,
		&s.ExpiresAt,
		&s.Actived,
	)

	if err == sql.ErrNoRows {
		return nil, session.ErrSessionNotFound
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Update implements [session.Repository].
func (r *SessionRepository) Update(ctx context.Context, sess *session.Session) error {
	query := `
		UPDATE sessions SET
			expires_at = $1,
			last_used_at = NOW(),
			actived = $2,
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query,
		sess.ExpiresAt,
		sess.LastUsedAt,
		sess.Actived,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return session.ErrSessionNotFound
	}

	return nil
}

func NewSessionRepository(db DBTX) session.Repository {
	return &SessionRepository{db: db}
}
