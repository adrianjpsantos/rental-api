package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/token"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/security"
)

type AuthService struct {
	UserService  user.Service
	TokenService token.Service
}

func (a *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenClaims, err := a.TokenService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	return a.TokenService.GenerateAccessToken(refreshTokenClaims.Payload())
}

func (a *AuthService) Logout(ctx context.Context, userID string) error {
	panic("unimplemented")
}

func (a *AuthService) Authenticate(
	ctx context.Context,
	input authenticate.AuthenticateInput,
) (authenticate.AuthenticateOutput, error) {
	emptyOutput := authenticate.AuthenticateOutput{}

	userForAuth, err := a.UserService.
		GetUserForAuthentication(ctx, input.Email)

	if err != nil {
		return emptyOutput, authenticate.ErrInvalidCredentials
	}

	err = security.CheckPassword(
		userForAuth.PasswordHash,
		input.Password,
	)

	if err != nil {
		return emptyOutput, authenticate.ErrInvalidCredentials
	}

	payload := authenticate.AuthenticatePayload{
		UserID: userForAuth.UserID,
	}

	accessToken, err := a.TokenService.GenerateAccessToken(payload)

	if err != nil {
		return emptyOutput, err
	}

	refreshToken, err := a.TokenService.GenerateRefreshToken(payload)

	if err != nil {
		return emptyOutput, err
	}

	return authenticate.AuthenticateOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func NewAuthService(userService user.Service, tokenService token.Service) authenticate.Service {
	return &AuthService{
		UserService:  userService,
		TokenService: tokenService,
	}
}
