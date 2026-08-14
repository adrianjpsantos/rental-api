package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, role Role) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error // soft delete
}

type Service interface {
	Create(ctx context.Context, role Role) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, id uuid.UUID, input User) error
	Delete(ctx context.Context, id uuid.UUID) error // soft delete
}
