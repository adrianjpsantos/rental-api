package token

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
)

type Service interface {
	GenerateAccessToken(ctx context.Context, payload authenticate.AuthenticatePayload) (string, error)
	GenerateRefreshToken(ctx context.Context, payload authenticate.AuthenticatePayload) (string, error)

	ValidateAccessToken(ctx context.Context, accessToken string) (TokenClaims, error)
	ValidateRefreshToken(ctx context.Context, refreshToken string) (TokenClaims, error)

	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
}
