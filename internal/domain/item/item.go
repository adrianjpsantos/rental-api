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

type ItemCreateInput struct {
	OwnerID     uuid.UUID `json:"owner_id"`
	CategoryID  uuid.UUID `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PricePerDay float64   `json:"price_per_day"`
	Quantity    int       `json:"quantity"`
}

type ItemUpdateInput struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	PricePerDay  float64  `json:"price_per_day"`
	PricePerHour float64  `json:"price_per_hour"`
	Quantity     int      `json:"quantity"`
	Photos       []string `json:"photos,omitempty"`
	Rules        string   `json:"rules,omitempty"`
}

type ItemFilter struct {
	OwnerID    *uuid.UUID
	CategoryID *uuid.UUID
	MinPrice   *float64
	MaxPrice   *float64
	Location   *string
	Limit      *int
	Offset     *int
}

// NewItem é o construtor da entidade
func NewItem(createInput ItemCreateInput) (*Item, error) {
	item := &Item{
		OwnerID:     createInput.OwnerID,
		CategoryID:  createInput.CategoryID,
		Title:       createInput.Title,
		Description: createInput.Description,
		PricePerDay: createInput.PricePerDay,
		Quantity:    createInput.Quantity,
		IsActive:    false, // por padrão, o item começa inativo ate que o usuário complete o cadastro
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := item.Validate(); err != nil {
		return nil, err
	}

	return item, nil
}

// Update atualiza os dados do item
func (i *Item) Update(update ItemUpdateInput) error {
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

func (i *Item) Validate() error {
	if i.OwnerID == uuid.Nil {
		return ErrInvalidOwnerID
	}
	if i.CategoryID == uuid.Nil {
		return ErrInvalidCategoryID
	}
	if len(i.Title) < 5 || len(i.Title) > 100 {
		return ErrInvalidTitle
	}
	if i.PricePerDay <= 0 {
		return ErrPricePerDayTooLow
	}
	if i.Quantity <= 0 {
		return ErrQuantityZeroOrNegative
	}

	// Valida condição
	validConditions := map[string]bool{
		"Novo": true, "Usado": true, "Semi-novo": true, "Reformado": true,
	}
	if !validConditions[i.Condition] {
		return ErrInvalidCondition
	}

	return nil
}

// IsAvailable verifica se o item tem quantidade disponível
func (i *Item) IsAvailableToRent(quantityRequested int) bool {
	return i.IsActive && i.Quantity >= quantityRequested
}

// CanBeRented verifica regras básicas para aluguel
func (i *Item) CanBeRented() error {
	if !i.IsActive {
		return ErrItemNotActive
	}
	if i.Quantity <= 0 {
		return ErrItemOutOfStock
	}
	return nil
}
