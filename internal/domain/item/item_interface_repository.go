package item

import (
	"context"

	"github.com/google/uuid"
)

// ItemRepository define o contrato para acesso aos dados
type InterfaceItemRepository interface {
	Create(ctx context.Context, item *Item) error
	GetByID(ctx context.Context, id uuid.UUID) (*Item, error)
	ListByFilters(ctx context.Context, filters ItemFilter) ([]*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

// ItemFilter é usado para filtros na listagem
type ItemFilter struct {
	OwnerID    *uuid.UUID
	CategoryID *uuid.UUID
	MinPrice   *float64
	MaxPrice   *float64
	Location   string
	Limit      int
	Offset     int
}
