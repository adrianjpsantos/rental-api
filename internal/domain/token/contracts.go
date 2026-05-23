package token

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
)

type Service interface {
	GenerateAccessToken(payload authenticate.AuthenticatePayload) (string, error)
	GenerateRefreshToken(payload authenticate.AuthenticatePayload) (string, error)

	ValidateAccessToken(accessToken string) (Claims, error)
	ValidateRefreshToken(refreshToken string) (Claims, error)
}
