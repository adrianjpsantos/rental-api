package application

import (
	"context"

	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/adrianjpsantos/rental-api/internal/security"
	"github.com/google/uuid"
)

type AuthAccountService struct {
	Repository authaccount.Repository
}

// Create implements [authaccount.Service].
func (a *AuthAccountService) Create(ctx context.Context, input authaccount.CreateInput) (*uuid.UUID, error) {

	accounts, err := a.Repository.ListByUserID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	isPrimary := len(accounts) == 0
	input.IsPrimary = isPrimary

	if input.PasswordHash != nil {
		passwordHashed, err := security.GenerateHashedPassword(*input.PasswordHash)
		if err != nil {
			return nil, err
		}
		input.PasswordHash = &passwordHashed
	}

	id, err := a.Repository.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// Delete implements [authaccount.Service].
func (a *AuthAccountService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetByID implements [authaccount.Service].
func (a *AuthAccountService) GetByID(ctx context.Context, id uuid.UUID) (*authaccount.AuthAccount, error) {
	panic("unimplemented")
}

// GetByProvider implements [authaccount.Service].
func (a *AuthAccountService) GetByProvider(ctx context.Context, provider authaccount.Provider, providerUserID string) (*authaccount.AuthAccount, error) {
	panic("unimplemented")
}

// GetLocalByEmail implements [authaccount.Service].
func (a *AuthAccountService) GetLocalByEmail(ctx context.Context, email string) (*authaccount.AuthAccount, error) {

	acc, err := a.Repository.FindLocalByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

// ListByUser implements [authaccount.Service].
func (a *AuthAccountService) ListByUser(ctx context.Context, userID uuid.UUID) ([]*authaccount.AuthAccount, error) {
	panic("unimplemented")
}

// Update implements [authaccount.Service].
func (a *AuthAccountService) Update(ctx context.Context, id uuid.UUID, input authaccount.UpdateInput) (*authaccount.AuthAccount, error) {
	panic("unimplemented")
}

func NewAuthAccountService(aaRepo authaccount.Repository) authaccount.Service {
	return &AuthAccountService{
		Repository: aaRepo,
	}
}
