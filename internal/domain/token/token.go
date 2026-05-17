package token

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	OIDToken *string `json:"oid_token,omitempty"`
	authenticate.AuthenticatePayload
	jwt.RegisteredClaims
}

func (t *TokenClaims) Payload() authenticate.AuthenticatePayload {
	return t.AuthenticatePayload
}
