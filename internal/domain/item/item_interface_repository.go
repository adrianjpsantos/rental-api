package item

import "github.com/google/uuid"

// ItemRepository define o contrato para acesso aos dados
type InterfaceItemRepository interface {
	Create(item *Item) error
	GetByID(id uuid.UUID) (*Item, error)
	ListByFilters(filters ItemFilter) ([]*Item, error)
	Update(item *Item) error
	Delete(id uuid.UUID) error
	Exists(id uuid.UUID) (bool, error)
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
