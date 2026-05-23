package application

import (
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/token"
	"github.com/adrianjpsantos/rental-api/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg *config.Config
}

func (s *TokenService) GenerateAccessToken(payload authenticate.AuthenticatePayload) (string, error) {

	duration, err := time.ParseDuration(s.cfg.JWT.AccessExpires)
	if err != nil {
		return "", err
	}

	claims := token.Claims{
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

func (s *TokenService) GenerateRefreshToken(payload authenticate.AuthenticatePayload) (string, error) {
	duration, err := time.ParseDuration(s.cfg.JWT.RefreshExpires)
	if err != nil {
		return "", err
	}

	claims := token.Claims{
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

func (s *TokenService) ValidateAccessToken(accessToken string) (token.Claims, error) {
	validToken, err := jwt.ParseWithClaims(accessToken, token.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.AccessSecret), nil
	})

	if err != nil {
		return token.Claims{}, err
	}

	return *validToken.Claims.(*token.Claims), nil
}

func (s *TokenService) ValidateRefreshToken(refreshToken string) (token.Claims, error) {
	validToken, err := jwt.ParseWithClaims(refreshToken, token.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.RefreshSecret), nil
	})

	if err != nil {
		return token.Claims{}, err
	}

	return *validToken.Claims.(*token.Claims), nil
}

func NewTokenService() token.Service {
	return &TokenService{
		cfg: config.LoadConfig(),
	}
}
