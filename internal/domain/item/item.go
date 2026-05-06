package item

import (
	"os/user"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/google/uuid"
)

type Item struct {
	Id            uuid.UUID
	OwnerID       uuid.UUID
	CategoryID    uuid.UUID
	Title         string
	Description   string
	Brand         string
	Model         string
	Year          int
	Condition     string
	PricePerDay   float64
	PricePerHour  float64
	MinRentalDays int
	MaxRentalDays int
	Quantity      int
	Location      string
	Latitude      float64
	Longitude     float64
	Photos        []string
	Rules         string
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Relacionamentos (opcional, mas útil)
	Owner    user.User
	Category category.Category // se tiver
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
