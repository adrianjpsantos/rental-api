package item

import (
	"context"

	"github.com/google/uuid"
)

// ItemRepository define o contrato para acesso aos dados
type Repository interface {
	Create(ctx context.Context, input *Item) error
	GetByID(ctx context.Context, id uuid.UUID) (*Item, error)
	ListByFilters(ctx context.Context, filters ItemFilter) ([]*Item, error)
	Update(ctx context.Context, update *Item) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

type Service interface {
	Create(ctx context.Context, input ItemCreateInput) error
	GetByID(ctx context.Context, id uuid.UUID) (*Item, error)
	ListByFilters(ctx context.Context, filters ItemFilter) ([]*Item, error)
	Update(ctx context.Context, id uuid.UUID, update ItemUpdateInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}
