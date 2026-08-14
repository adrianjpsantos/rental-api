package repositories

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	db DBTX
}

func NewCategoryRepository(db DBTX) category.Repository {
	return &CategoryRepository{db: db}
}

// Create
func (r *CategoryRepository) Create(ctx context.Context, input category.Category) error {
	query := `
		INSERT INTO categories (
			id, name, slug, description,
			is_active, icon, position,
			created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := r.db.ExecContext(ctx, query,
		input.ID,
		input.Name,
		input.Slug,
		input.Description,
		input.IsActive,
		input.Icon,
		input.Position,
		input.CreatedAt,
		input.UpdatedAt,
	)

	return err
}

// GetByID
func (r *CategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*category.Category, error) {
	query := `
		SELECT id, name, slug, description,
		       is_active, icon, position,
		       created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	var c category.Category

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.Name,
		&c.Slug,
		&c.Description,
		&c.IsActive,
		&c.Icon,
		&c.Position,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, category.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// List
func (r *CategoryRepository) List(ctx context.Context) ([]*category.Category, error) {
	query := `
		SELECT id, name, slug, description,
		       is_active, icon, position,
		       created_at, updated_at
		FROM categories
		WHERE is_active = true
		ORDER BY position ASC, name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*category.Category

	for rows.Next() {
		var c category.Category

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.Description,
			&c.IsActive,
			&c.Icon,
			&c.Position,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &c)
	}

	if len(list) == 0 {
		return nil, category.ErrCategoryNotFound
	}

	return list, nil
}

// Update
func (r *CategoryRepository) Update(ctx context.Context, update category.Category) error {
	query := `
		UPDATE categories SET
			name = $1,
			slug = $2,
			description = $3,
			is_active = $4,
			icon = $5,
			position = $6,
			updated_at = NOW()
		WHERE id = $7
	`

	result, err := r.db.ExecContext(ctx, query,
		update.Name,
		update.Slug,
		update.Description,
		update.IsActive,
		update.Icon,
		update.Position,
		update.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return category.ErrCategoryNotFound
	}

	return nil
}

// Delete (soft delete)
func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE categories
		SET is_active = false,
		    updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return category.ErrCategoryNotFound
	}

	return nil
}

// Exists (por nome)
func (r *CategoryRepository) Exists(ctx context.Context, name string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM categories
			WHERE LOWER(name) = LOWER($1)
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
