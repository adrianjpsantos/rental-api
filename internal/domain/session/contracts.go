package session

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
)

type Service interface {
	GenerateAccessToken(payload authenticate.AuthenticatePayload) (string, error)
	GenerateRefreshToken(payload authenticate.AuthenticatePayload) (string, error)
	ValidateAccessToken(accessToken string) (*Claims, error)
	ValidateRefreshToken(refreshToken string) (*Claims, error)
	DesactivateSession(ctx context.Context, sessionId string) error
}

type Repository interface {
	Create(ctx context.Context, session *Session) error
	FindById(ctx context.Context, id string) (*Session, error)
	Update(ctx context.Context, session *Session) error
	Delete(ctx context.Context, id string) error
	Desactive(ctx context.Context, id string) error
}
