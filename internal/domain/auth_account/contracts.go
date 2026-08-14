package authaccount

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, account *CreateInput) (*uuid.UUID, error)
	Update(ctx context.Context, account *AuthAccount) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*AuthAccount, error)

	FindLocalByEmail(ctx context.Context, email string) (*AuthAccount, error)

	FindByProvider(
		ctx context.Context,
		provider Provider,
		providerUserID string,
	) (*AuthAccount, error)

	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*AuthAccount, error)
}

type Service interface {
	Create(ctx context.Context, input CreateInput) (*uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*AuthAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error

	GetByID(ctx context.Context, id uuid.UUID) (*AuthAccount, error)

	GetLocalByEmail(ctx context.Context, email string) (*AuthAccount, error)

	GetByProvider(
		ctx context.Context,
		provider Provider,
		providerUserID string,
	) (*AuthAccount, error)

	ListByUser(ctx context.Context, userID uuid.UUID) ([]*AuthAccount, error)
}
