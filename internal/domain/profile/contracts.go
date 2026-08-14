package profile

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, profile *CreateInput) error
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, userID uuid.UUID) error

	FindByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error)
}

type Service interface {
	Create(ctx context.Context, input CreateInput) error
	Update(ctx context.Context, userID uuid.UUID, input UpdateInput) (*Profile, error)
	Delete(ctx context.Context, userID uuid.UUID) error

	GetByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error)
}
