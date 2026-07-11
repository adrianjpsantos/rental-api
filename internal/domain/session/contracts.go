package session

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GenerateAccessToken(ctx context.Context, userId uuid.UUID) (string, error)
	GenerateRefreshToken(ctx context.Context, userId uuid.UUID, sessionId uuid.UUID) (string, *time.Time, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (*Claims, error)
	ValidateRefreshToken(ctx context.Context, refreshToken string) (*Claims, error)
	StartSession(ctx context.Context, userId uuid.UUID) (string, string, error)
	RefreshSession(ctx context.Context, refreshToken string) (string, error)
	DesactivateSession(ctx context.Context, sessionId string) error
}

type Repository interface {
	Create(ctx context.Context, session *Session) error
	FindByHash(ctx context.Context, hash string) (*Session, error)
	Update(ctx context.Context, session *Session) error
	Delete(ctx context.Context, hash string) error
	Desactive(ctx context.Context, hash string) error
}
