package authenticate

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
)

type Service interface {
	Authenticate(ctx context.Context, authenticateInput AuthenticateInput) (*AuthenticateOutput, error)
	Logout(ctx context.Context, userID string) error
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
	SignUp(ctx context.Context, input user.UserCreateInput) (*AuthenticateOutput, error)
}
