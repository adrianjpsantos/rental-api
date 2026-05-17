package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/authenticate"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
)

type AuthService struct {
	userRepo user.InterfaceUserRepository
}

func (a *AuthService) Authenticate(ctx context.Context, authenticateInput authenticate.AuthenticateInput) (user.User, error) {
	panic("unimplemented")
}

func NewAuthService(userRepo user.InterfaceUserRepository) authenticate.InterfaceAuthenticateService {
	return &AuthService{
		userRepo: userRepo,
	}
}
