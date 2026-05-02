package item

import (
	"os/user"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Id            uuid.UUID `gorm:"type:uuid;primaryKey"`
	OwnerID       uuid.UUID `gorm:"type:uuid;not null;index"`
	CategoryID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Title         string    `gorm:"not null;size:255"`
	Description   string    `gorm:"type:text"`
	Brand         string    `gorm:"size:100"`
	Model         string    `gorm:"size:100"`
	Year          int       `gorm:"index"`
	Condition     string    `gorm:"size:50"`
	PricePerDay   float64   `gorm:"not null;type:decimal(10,2)"`
	PricePerHour  float64   `gorm:"type:decimal(10,2)"`
	MinRentalDays int       `gorm:"default:1"`
	MaxRentalDays int       `gorm:"default:30"`
	Quantity      int       `gorm:"not null;default:1"`
	Location      string    `gorm:"size:255"`
	Latitude      float64   `gorm:"type:decimal(10,8)"`
	Longitude     float64   `gorm:"type:decimal(11,8)"`
	Photos        []string  `gorm:"type:text[]"` // Para PostgreSQL
	// Photos     string    `gorm:"type:text"`             // Alternativa (JSON)
	Rules     string `gorm:"type:text"`
	IsActive  bool   `gorm:"not null;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relacionamentos (opcional, mas útil)
	Owner    user.User         `gorm:"foreignKey:OwnerID"`
	Category category.Category `gorm:"foreignKey:CategoryID"` // se tiver
}

type ItemUpdate struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	PricePerDay  float64  `json:"price_per_day"`
	PricePerHour float64  `json:"price_per_hour"`
	Quantity     int      `json:"quantity"`
	Photos       []string `json:"photos,omitempty"`
	Rules        string   `json:"rules,omitempty"`
}

// NewItem é o construtor da entidade
func NewItem(ownerID, categoryID uuid.UUID, title, description string, pricePerDay float64, quantity int) (*Item, error) {
	item := &Item{
		Id:          uuid.New(),
		OwnerID:     ownerID,
		CategoryID:  categoryID,
		Title:       title,
		Description: description,
		PricePerDay: pricePerDay,
		Quantity:    quantity,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := item.Validate(); err != nil {
		return nil, err
	}

	return item, nil
}

// Update atualiza os dados do item
func (i *Item) Update(update ItemUpdate) error {
	updated := false

	if update.Title != "" {
		i.Title = update.Title
		updated = true
	}
	if update.Description != "" {
		i.Description = update.Description
		updated = true
	}
	if update.PricePerDay > 0 {
		i.PricePerDay = update.PricePerDay
		updated = true
	}
	if update.Quantity > 0 {
		i.Quantity = update.Quantity
		updated = true
	}
	if len(update.Photos) > 0 {
		i.Photos = update.Photos
		updated = true
	}
	if update.Rules != "" {
		i.Rules = update.Rules
		updated = true
	}

	if updated {
		i.UpdatedAt = time.Now()
	}

	return i.Validate()
}
