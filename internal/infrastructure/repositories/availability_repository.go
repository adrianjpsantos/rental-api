package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/google/uuid"
)

type AvailabilityRepository struct {
	db DBTX
}

func NewAvailabilityRepository(db DBTX) availability.Repository {
	return &AvailabilityRepository{db: db}
}

// Create
func (r *AvailabilityRepository) Create(ctx context.Context, slot availability.AvailabilitySlot) error {
	query := `
		INSERT INTO availability_slots (
			id, item_id, start_date, end_date,
			type, reason, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.db.ExecContext(ctx, query,
		slot.Id,
		slot.ItemId,
		slot.StartDate,
		slot.EndDate,
		slot.Type,
		slot.Reason,
		slot.CreatedAt,
	)

	return err
}

// GetByID
func (r *AvailabilityRepository) GetByID(ctx context.Context, id uuid.UUID) (*availability.AvailabilitySlot, error) {
	query := `
		SELECT id, item_id, start_date, end_date,
		       type, reason, created_at
		FROM availability_slots
		WHERE id = $1
	`

	var slot availability.AvailabilitySlot

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&slot.Id,
		&slot.ItemId,
		&slot.StartDate,
		&slot.EndDate,
		&slot.Type,
		&slot.Reason,
		&slot.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, availability.ErrSlotNotFound
	}
	if err != nil {
		return nil, err
	}

	return &slot, nil
}

// FindByItemID
func (r *AvailabilityRepository) ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*availability.AvailabilitySlot, error) {
	query := `
		SELECT id, item_id, start_date, end_date,
		       type, reason, created_at
		FROM availability_slots
		WHERE item_id = $1
		ORDER BY start_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*availability.AvailabilitySlot

	for rows.Next() {
		var slot availability.AvailabilitySlot

		err := rows.Scan(
			&slot.Id,
			&slot.ItemId,
			&slot.StartDate,
			&slot.EndDate,
			&slot.Type,
			&slot.Reason,
			&slot.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &slot)
	}

	if len(list) == 0 {
		return nil, availability.ErrSlotNotFound
	}

	return list, nil
}

// FindOverlapping (REGRA CRÍTICA)
func (r *AvailabilityRepository) ListOverlapping(
	ctx context.Context,
	itemID uuid.UUID,
	startDate, endDate time.Time,
) ([]*availability.AvailabilitySlot, error) {

	query := `
		SELECT id, item_id, start_date, end_date,
		       type, reason, created_at
		FROM availability_slots
		WHERE item_id = $1
		  AND start_date < $3
		  AND end_date > $2
	`

	rows, err := r.db.QueryContext(ctx, query, itemID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*availability.AvailabilitySlot

	for rows.Next() {
		var slot availability.AvailabilitySlot

		err := rows.Scan(
			&slot.Id,
			&slot.ItemId,
			&slot.StartDate,
			&slot.EndDate,
			&slot.Type,
			&slot.Reason,
			&slot.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &slot)
	}

	return list, nil
}

// Delete
func (r *AvailabilityRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM availability_slots WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return availability.ErrSlotNotFound
	}

	return nil
}

// HasBlockingSlot (rápido e eficiente)
func (r *AvailabilityRepository) ExistsBlockingSlot(
	ctx context.Context,
	itemID uuid.UUID,
	startDate, endDate time.Time,
) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1 FROM availability_slots
			WHERE item_id = $1
			  AND start_date < $3
			  AND end_date > $2
			  AND type = 'blocked'
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, itemID, startDate, endDate).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
