package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/token"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg *config.Config
}

func (s *TokenService) GenerateAccessToken(ctx context.Context, payload authenticate.AuthenticatePayload) (string, error) {

	duration, err := time.ParseDuration(s.cfg.JWT.AccessExpires)
	if err != nil {
		return "", err
	}

	claims := token.TokenClaims{
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

func (s *TokenService) GenerateRefreshToken(ctx context.Context, payload authenticate.AuthenticatePayload) (string, error) {
	duration, err := time.ParseDuration(s.cfg.JWT.RefreshExpires)
	if err != nil {
		return "", err
	}

	claims := token.TokenClaims{
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

func (s *TokenService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenClaims, err := s.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	return s.GenerateAccessToken(ctx, refreshTokenClaims.Payload())
}

func (s *TokenService) ValidateAccessToken(ctx context.Context, accessToken string) (token.TokenClaims, error) {
	validToken, err := jwt.ParseWithClaims(accessToken, token.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.AccessSecret), nil
	})

	if err != nil {
		return token.TokenClaims{}, err
	}

	return *validToken.Claims.(*token.TokenClaims), nil
}

func (s *TokenService) ValidateRefreshToken(ctx context.Context, refreshToken string) (token.TokenClaims, error) {
	validToken, err := jwt.ParseWithClaims(refreshToken, token.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.RefreshSecret), nil
	})

	if err != nil {
		return token.TokenClaims{}, err
	}

	return *validToken.Claims.(*token.TokenClaims), nil
}

func NewTokenService() token.InterfaceTokenService {
	return &TokenService{
		cfg: config.LoadConfig(),
	}
}
