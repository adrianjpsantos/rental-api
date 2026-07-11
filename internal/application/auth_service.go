package application

import (
	"context"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/session"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/security"
)

type AuthService struct {
	UserService    user.Service
	SessionService session.Service
}

func (a *AuthService) SignUp(ctx context.Context, input user.UserCreateInput) (*authenticate.AuthenticateOutput, error) {
	err := a.UserService.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	authenticated, err := a.Authenticate(ctx, authenticate.AuthenticateInput{Email: input.Email,
		Password: input.Password})

	return authenticated, nil
}

func (a *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	return a.SessionService.RefreshSession(ctx, refreshToken)
}

func (a *AuthService) Logout(ctx context.Context, userID string) error {
	panic("unimplemented")
}

func (a *AuthService) Authenticate(
	ctx context.Context,
	input authenticate.AuthenticateInput,
) (*authenticate.AuthenticateOutput, error) {

	userForAuth, err := a.UserService.
		GetUserForAuthentication(ctx, input.Email)

	if err != nil {
		return nil, authenticate.ErrInvalidCredentials
	}

	err = security.CheckPassword(
		userForAuth.PasswordHash,
		input.Password,
	)

	if err != nil {
		return nil, authenticate.ErrInvalidCredentials
	}

	refreshToken, accessToken, err := a.SessionService.StartSession(ctx, userForAuth.UserID)

	if err != nil {
		fmt.Println("Authenticate: ", err)
		return nil, err

	}

	return &authenticate.AuthenticateOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func NewAuthService(userService user.Service, sessionService session.Service) authenticate.Service {
	return &AuthService{
		UserService:    userService,
		SessionService: sessionService,
	}
}
