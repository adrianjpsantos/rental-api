package authenticate

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
)

type InterfaceAuthenticateService interface {
	Authenticate(ctx context.Context, authenticateInput AuthenticateInput) (user.User, error)
}
