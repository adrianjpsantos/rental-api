package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type SessionService struct {
	cfg        *config.Config
	repository session.Repository
}

// DesactivateSession implements [session.Service].
func (s *SessionService) DesactivateSession(ctx context.Context, id string) error {
	return s.repository.Desactive(ctx, id)
}

func (s *SessionService) GenerateAccessToken(payload authenticate.AuthenticatePayload) (string, error) {

	duration, err := time.ParseDuration(s.cfg.JWT.AccessExpires)
	if err != nil {
		return "", err
	}

	claims := session.Claims{
		AuthenticatePayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(duration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := s.cfg.JWT.AccessSecret

	return generateToken.SignedString([]byte(secretKey))
}

func (s *SessionService) GenerateRefreshToken(payload authenticate.AuthenticatePayload) (string, error) {
	duration, err := time.ParseDuration(s.cfg.JWT.RefreshExpires)
	if err != nil {
		return "", err
	}

	claims := session.Claims{
		AuthenticatePayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(duration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := s.cfg.JWT.RefreshSecret

	return generateToken.SignedString([]byte(secretKey))
}

func (s *SessionService) ValidateAccessToken(accessToken string) (*session.Claims, error) {
	validToken, err := jwt.ParseWithClaims(accessToken, session.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return validToken.Claims.(*session.Claims), nil
}

func (s *SessionService) ValidateRefreshToken(refreshToken string) (*session.Claims, error) {
	validToken, err := jwt.ParseWithClaims(refreshToken, session.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return validToken.Claims.(*session.Claims), nil
}

func NewSessionService(rep session.Repository) session.Service {
	return &SessionService{
		cfg:        config.LoadConfig(),
		repository: rep,
	}
}
