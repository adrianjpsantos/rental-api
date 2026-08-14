package authenticate

import (
	"context"
)

type Service interface {
	LoginLocal(ctx context.Context, authenticateInput AuthenticateInput) (*AuthenticateOutput, error)
	Logout(ctx context.Context, userID string) error
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
	Register(ctx context.Context, input RegisterInput) error
}
