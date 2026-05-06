package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/google/uuid"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) item.InterfaceItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(ctx context.Context, item *item.Item) error {
	query := `
		INSERT INTO items (
			id, owner_id, category_id, title, description,
			brand, model, year, condition,
			price_per_day, price_per_hour,
			min_rental_days, max_rental_days,
			quantity, location, latitude, longitude,
			photos, rules, is_active,
			created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,
			$10,$11,$12,$13,
			$14,$15,$16,$17,
			$18,$19,$20,
			$21,$22
		)
	`

	photosJSON, _ := json.Marshal(item.Photos)

	_, err := r.db.ExecContext(ctx, query,
		item.Id,
		item.OwnerID,
		item.CategoryID,
		item.Title,
		item.Description,
		item.Brand,
		item.Model,
		item.Year,
		item.Condition,
		item.PricePerDay,
		item.PricePerHour,
		item.MinRentalDays,
		item.MaxRentalDays,
		item.Quantity,
		item.Location,
		item.Latitude,
		item.Longitude,
		photosJSON,
		item.Rules,
		item.IsActive,
		item.CreatedAt,
		item.UpdatedAt,
	)

	return err
}

func (r *ItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*item.Item, error) {
	query := `
		SELECT id, owner_id, category_id, title, description,
		       brand, model, year, condition,
		       price_per_day, price_per_hour,
		       min_rental_days, max_rental_days,
		       quantity, location, latitude, longitude,
		       photos, rules, is_active,
		       created_at, updated_at
		FROM items
		WHERE id = $1
	`

	var it item.Item
	var photos []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&it.Id,
		&it.OwnerID,
		&it.CategoryID,
		&it.Title,
		&it.Description,
		&it.Brand,
		&it.Model,
		&it.Year,
		&it.Condition,
		&it.PricePerDay,
		&it.PricePerHour,
		&it.MinRentalDays,
		&it.MaxRentalDays,
		&it.Quantity,
		&it.Location,
		&it.Latitude,
		&it.Longitude,
		&photos,
		&it.Rules,
		&it.IsActive,
		&it.CreatedAt,
		&it.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, item.ErrItemNotFound
	}
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(photos, &it.Photos)

	return &it, nil
}
func (r *ItemRepository) ListByFilters(ctx context.Context, f item.ItemFilter) ([]*item.Item, error) {
	query := `
		SELECT id, owner_id, category_id, title, description,
		       brand, model, year, condition,
		       price_per_day, price_per_hour,
		       min_rental_days, max_rental_days,
		       quantity, location, latitude, longitude,
		       photos, rules, is_active,
		       created_at, updated_at
		FROM items
		WHERE is_active = true
	`

	args := []interface{}{}
	argPos := 1

	if f.OwnerID != nil {
		query += fmt.Sprintf(" AND owner_id = $%d", argPos)
		args = append(args, *f.OwnerID)
		argPos++
	}

	if f.CategoryID != nil {
		query += fmt.Sprintf(" AND category_id = $%d", argPos)
		args = append(args, *f.CategoryID)
		argPos++
	}

	if f.MinPrice != nil {
		query += fmt.Sprintf(" AND price_per_day >= $%d", argPos)
		args = append(args, *f.MinPrice)
		argPos++
	}

	if f.MaxPrice != nil {
		query += fmt.Sprintf(" AND price_per_day <= $%d", argPos)
		args = append(args, *f.MaxPrice)
		argPos++
	}

	if f.Location != "" {
		query += fmt.Sprintf(" AND location ILIKE $%d", argPos)
		args = append(args, "%"+f.Location+"%")
		argPos++
	}

	query += " ORDER BY created_at DESC"

	if f.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argPos)
		args = append(args, f.Limit)
		argPos++
	}

	if f.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argPos)
		args = append(args, f.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*item.Item

	for rows.Next() {
		var it item.Item
		var photos []byte

		err := rows.Scan(
			&it.Id,
			&it.OwnerID,
			&it.CategoryID,
			&it.Title,
			&it.Description,
			&it.Brand,
			&it.Model,
			&it.Year,
			&it.Condition,
			&it.PricePerDay,
			&it.PricePerHour,
			&it.MinRentalDays,
			&it.MaxRentalDays,
			&it.Quantity,
			&it.Location,
			&it.Latitude,
			&it.Longitude,
			&photos,
			&it.Rules,
			&it.IsActive,
			&it.CreatedAt,
			&it.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		_ = json.Unmarshal(photos, &it.Photos)

		list = append(list, &it)
	}

	if len(list) == 0 {
		return nil, item.ErrItemNotFound
	}

	return list, nil
}
func (r *ItemRepository) Update(ctx context.Context, i *item.Item) error {
	query := `
		UPDATE items SET
			title = $1,
			description = $2,
			price_per_day = $3,
			price_per_hour = $4,
			quantity = $5,
			photos = $6,
			rules = $7,
			updated_at = NOW()
		WHERE id = $8
	`

	photosJSON, _ := json.Marshal(i.Photos)

	result, err := r.db.ExecContext(ctx, query,
		i.Title,
		i.Description,
		i.PricePerDay,
		i.PricePerHour,
		i.Quantity,
		photosJSON,
		i.Rules,
		i.Id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return item.ErrItemNotFound
	}

	return nil
}
func (r *ItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE items
		SET is_active = false,
		    updated_at = NOW()
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return item.ErrItemNotFound
	}

	return nil
}
func (r *ItemRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM items WHERE id = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)

	return exists, err
}
