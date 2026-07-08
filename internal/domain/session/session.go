package session

import (
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID `json:"id,omitempty" validate:"required,uuid"`
	UserId    uuid.UUID `json:"user_id,omitempty" validate:"required,uuid"`
	Token     string    `json:"token,omitempty" validate:"required"`
	ExpiresAt time.Time `json:"expires_at,omitempty" validate:"datetime"`
	CreatedAt time.Time `json:"created_at,omitempty" validate:"datetime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" validate:"datetime"`
	Actived   bool      `json:"actived,omitempty" validate:"required,boolean"`
}

type Claims struct {
	OIDToken *string `json:"oid_token,omitempty"`
	authenticate.AuthenticatePayload
	jwt.RegisteredClaims
}

func (t *Claims) Payload() authenticate.AuthenticatePayload {
	return t.AuthenticatePayload
}
