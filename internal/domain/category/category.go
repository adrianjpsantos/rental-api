package category

import (
	"regexp"
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

type CategoryCreateInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Position    int    `json:"position" validate:"required"`
}

func NewCategory(newCat CategoryCreateInput) (*Category, error) {
	category := &Category{
		ID:          uuid.New(),
		Name:        strings.TrimSpace(newCat.Name),
		Description: strings.TrimSpace(newCat.Description),
		Icon:        strings.TrimSpace(newCat.Icon),
		Position:    newCat.Position,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	category.GenerateSlug()

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
		c.GenerateSlug() // ← importante!
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

func (c *Category) Validate() error {
	if err := c.validateName(); err != nil {
		return err
	}
	if err := c.validateSlug(); err != nil {
		return err
	}
	if err := c.validateDescription(); err != nil {
		return err
	}
	return nil
}

func (c *Category) validateName() error {
	name := strings.TrimSpace(c.Name)
	if name == "" {
		return ErrInvalidName
	}
	if len(name) < 3 {
		return ErrNameTooShort
	}
	if len(name) > 80 {
		return ErrNameTooLong
	}
	return nil
}

func (c *Category) validateSlug() error {
	if c.Slug == "" {
		return ErrInvalidSlug
	}
	if !c.IsValidSlug() {
		return ErrSlugInvalidFormat
	}
	return nil
}

func (c *Category) validateDescription() error {
	if len(strings.TrimSpace(c.Description)) > 500 {
		return ErrDescriptionTooLong
	}
	return nil
}

// generateSlug gera um slug a partir do nome
func (c *Category) GenerateSlug() string {
	slug := strings.ToLower(strings.TrimSpace(c.Name))
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove caracteres especiais
	slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")
	// Remove hífens duplicados
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

// isValidSlug verifica se o slug está no formato correto
func (c *Category) IsValidSlug() bool {
	match, _ := regexp.MatchString(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, c.Slug)
	return match
}
