package category

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, input Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, name string) (bool, error)
}

type Service interface {
	Create(ctx context.Context, input CategoryCreateInput) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, id uuid.UUID, update CategoryUpdate) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, name string) (bool, error)
}
