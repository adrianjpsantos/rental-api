package category

import (
	"context"

	"github.com/google/uuid"
)

type InterfaceCategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, name string) (bool, error)
}
