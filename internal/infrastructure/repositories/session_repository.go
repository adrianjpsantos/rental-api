package repositories

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
)

type SessionRepository struct {
	db *sql.DB
}

// Desactive implements [session.Repository].
func (r *SessionRepository) Desactive(ctx context.Context, token string) error {
	query := `
		UPDATE sessions SET
			actived = false,
			updated_at = NOW()
		WHERE token = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		token,
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
		(id, user_id, token, expires_at,created_at,updated_at ) 
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.ExecContext(ctx, query,
		session.Id,
		session.UserId,
		session.Token,
		session.ExpiresAt,
		session.CreatedAt,
		session.UpdatedAt,
	)

	return err
}

// Delete implements [session.Repository].
func (r *SessionRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindById implements [session.Repository].
func (r *SessionRepository) FindById(ctx context.Context, id string) (*session.Session, error) {
	panic("unimplemented")
}

// Update implements [session.Repository].
func (r *SessionRepository) Update(ctx context.Context, session *session.Session) error {
	panic("unimplemented")
}

func NewSessionRepository(db *sql.DB) session.Repository {
	return &SessionRepository{db: db}
}
