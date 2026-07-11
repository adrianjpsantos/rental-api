package application

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SessionService struct {
	cfg        *config.Config
	repository session.Repository
}

// DesactivateSession implements [session.Service].
func (s *SessionService) DesactivateSession(ctx context.Context, id string) error {
	return s.repository.Desactive(ctx, id)
}

func (s *SessionService) RefreshSession(ctx context.Context, refreshToken string) (string, error) {
	claims, err := s.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := s.GenerateAccessToken(ctx, claims.UserId)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *SessionService) StartSession(ctx context.Context, userId uuid.UUID) (string, string, error) {
	sess := session.NewSession(userId)
	refreshToken, expiredAt, err := s.GenerateRefreshToken(ctx, userId, sess.Id)

	if err != nil {
		return "", "", err
	}

	tokenHash := session.FormatTokenHash(refreshToken)

	sess.UpdateToken(tokenHash, *expiredAt)

	err = s.repository.Create(ctx, sess)
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.GenerateAccessToken(ctx, userId)

	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (s *SessionService) GenerateAccessToken(ctx context.Context, userId uuid.UUID) (string, error) {

	duration, err := time.ParseDuration(s.cfg.JWT.AccessExpires)
	if err != nil {
		return "", err
	}

	claims := session.Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(duration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := s.cfg.JWT.AccessSecret

	signed, err := generateToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signed, err
}

func (s *SessionService) GenerateRefreshToken(ctx context.Context, userId uuid.UUID, sessionId uuid.UUID) (string, *time.Time, error) {
	duration, err := time.ParseDuration(s.cfg.JWT.RefreshExpires)
	if err != nil {
		return "", nil, err
	}

	claims := session.Claims{
		SessionId: sessionId,
		UserId:    userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(duration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := s.cfg.JWT.RefreshSecret

	signed, err := generateToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", nil, err
	}

	return signed, &claims.ExpiresAt.Time, err
}

func (s *SessionService) ValidateAccessToken(ctx context.Context, accessToken string) (*session.Claims, error) {
	claims := &session.Claims{}

	validToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.AccessSecret), nil
	})

	if err != nil {
		fmt.Printf("%T\n", err)
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	if !validToken.Valid {
		return nil, session.ErrInvalidToken
	}

	return claims, nil
}

func (s *SessionService) ValidateRefreshToken(ctx context.Context, refreshToken string) (*session.Claims, error) {
	tokenHash := session.FormatTokenHash(refreshToken)
	sess, err := s.repository.FindByHash(ctx, tokenHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, session.ErrSessionNotFound
		}
		return nil, err
	}

	if sess.Expired() {
		return nil, session.ErrSessionExpired
	}

	claims := &session.Claims{}

	validToken, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.RefreshSecret), nil
	})

	if err != nil {
		fmt.Printf("%T\n", err)
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	if !validToken.Valid {
		return nil, session.ErrInvalidRefreshToken
	}

	return claims, nil

}

func NewSessionService(rep session.Repository) session.Service {
	return &SessionService{
		cfg:        config.LoadConfig(),
		repository: rep,
	}
}
