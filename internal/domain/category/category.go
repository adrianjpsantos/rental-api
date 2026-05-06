package category

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Category representa uma categoria de itens no sistema
type Category struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	IsActive    bool
	Icon        string
	Position    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CategoryUpdate - DTO para atualização
type CategoryUpdate struct {
	Name        string
	Description string
	Icon        string
	Position    int
	IsActive    bool
}

func NewCategory(name, description, icon string, position int) (*Category, error) {
	category := &Category{
		ID:          uuid.New(),
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Icon:        strings.TrimSpace(icon),
		Position:    position,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	category.Slug = generateSlug(category.Name)

	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

// Update - Método atualizado
func (c *Category) Update(update CategoryUpdate) error {
	updated := false

	if update.Name != "" {
		c.Name = strings.TrimSpace(update.Name)
		c.Slug = generateSlug(c.Name) // ← importante!
		updated = true
	}

	if update.Description != "" {
		c.Description = strings.TrimSpace(update.Description)
		updated = true
	}

	if update.Icon != "" {
		c.Icon = strings.TrimSpace(update.Icon)
		updated = true
	}

	if update.Position != 0 && update.Position != c.Position {
		c.Position = update.Position
		updated = true
	}

	if update.IsActive != c.IsActive {
		c.IsActive = update.IsActive
		updated = true
	}

	if updated {
		c.UpdatedAt = time.Now()
		return c.Validate()
	}

	return nil
}
