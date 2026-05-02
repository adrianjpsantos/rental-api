package category

import (
	"github.com/google/uuid"
)

type Category struct {
	Id      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name    string    `gorm:"type:text"`
	IconURL string    `gorm:"type:text"`

	// Índice único para evitar múltiplas avaliações do mesmo rental
	UniqueRentalReview string `gorm:"-"` // apenas referência
}

func NewCategory(name, iconURL string) (*Category, error) {
	item := &Category{
		Id:      uuid.New(),
		Name:    name,
		IconURL: iconURL,
	}

	return item, nil
}
