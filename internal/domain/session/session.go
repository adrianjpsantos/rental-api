package session

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID `json:"id,omitempty" validate:"required,uuid"`
	UserId    uuid.UUID `json:"user_id,omitempty" validate:"required,uuid"`
	TokenHash string    `json:"token_hash,omitempty" validate:"required"`
	ExpiresAt time.Time `json:"expires_at,omitempty" validate:"datetime"`
	CreatedAt time.Time `json:"created_at,omitempty" validate:"datetime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" validate:"datetime"`
	Actived   bool      `json:"actived,omitempty" validate:"required,boolean"`
}

type Claims struct {
	SessionId uuid.UUID `json:"session_id,omitempty" validate:"required,uuid"`
	UserId    uuid.UUID `json:"user_id,omitempty" validate:"required,uuid"`
	jwt.RegisteredClaims
}

func (s *Session) Expired() bool {
	return s.ExpiresAt.Before(time.Now()) || !s.Actived
}

func NewSession(userId uuid.UUID) *Session {
	return &Session{
		Id:        uuid.New(),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Actived:   true,
	}
}

func FormatTokenHash(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *Session) UpdateToken(tokenHash string, expiresAt time.Time) {
	s.TokenHash = tokenHash
	s.ExpiresAt = expiresAt
	s.UpdatedAt = time.Now()
}
