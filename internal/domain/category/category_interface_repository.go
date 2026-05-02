package category

import (
	"github.com/google/uuid"
)

// ItemRepository define o contrato para acesso aos dados
type InterfaceCategoryRepository interface {
	Create(category *Category) error
	GetByID(id uuid.UUID) (*Category, error)
	Update(category *Category) error
	Delete(id uuid.UUID) error
	Exists(categoryName string) (bool, error)
}
