package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/google/uuid"
)

type RentalRepository struct {
	db *sql.DB
}

// Create implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) Create(ctx context.Context, rental *rental.Rental) error {

	query := `INSERT INTO rentals (id,item_id,lessee_id,lessor_id,start_date,end_date,total_amount,status,payment_status,delivery_method,notes,created_at,update_at,started_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`

	_, err := r.db.ExecContext(ctx, query,
		rental.Id,
		rental.ItemID,
		rental.LesseeID,
		rental.LessorID,
		rental.StartDate,
		rental.EndDate,
		rental.TotalAmount,
		rental.Status,
		rental.PaymentStatus,
		rental.DeliveryMethod,
		rental.Notes,
		rental.CreatedAt,
		rental.UpdatedAt,
		rental.StartedAt,
	)

	if err != nil {
		fmt.Printf("CREATE RENTAL ERR: %s", err.Error())
		return err
	}

	return nil
}

// ExistsOverlapping implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) ExistsOverlapping(ctx context.Context, itemID uuid.UUID, startDate time.Time, endDate time.Time) (bool, error) {

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM rental WHERE item_id = $1 AND start_date = $2 AND end_date = $3)`
	err := r.db.QueryRowContext(ctx, query, itemID, startDate, endDate).Scan(&exists)

	if err != nil {
		return false, err
	}

	fmt.Println("❕USER EXISTS BY CPF: ", exists)
	return exists, nil
}

// GetByID implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) GetByID(ctx context.Context, id uuid.UUID) (*rental.Rental, error) {
	var rent rental.Rental

	query := `SELECT id,item_id,lessee_id,lessor_id,start_date,end_date,total_amount,status,payment_status,delivery_method,notes,created_at,update_at,started_at,complete_at,cancelled_at
	FROM rentals WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rent.Id,
		&rent.ItemID,
		&rent.LesseeID,
		&rent.LessorID,
		&rent.StartDate,
		&rent.EndDate,
		&rent.TotalAmount,
		&rent.Status,
		&rent.PaymentStatus,
		&rent.DeliveryMethod,
		&rent.Notes,
		&rent.CreatedAt,
		&rent.UpdatedAt,
		&rent.StartedAt,
		&rent.CompletedAt,
		&rent.CancelledAt,
	)

	if err != nil {
		fmt.Printf("☠️RENTAL GET BY ID ERR: %s", err)
		return nil, err
	}

	return &rent, nil

}

// ListByLessee implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	var rentals []*rental.Rental

	query := `SELECT id, item_id, lessee_id, lessor_id, start_date, end_date, total_amount, status, payment_status, delivery_method, notes, created_at, updated_at, started_at, completed_at, cancelled_at
			  FROM rentals WHERE lessee_id = $1`

	args := []interface{}{lesseeID}

	if status != nil {
		query += " AND status = $2"
		args = append(args, *status)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rent rental.Rental

		err := rows.Scan(
			&rent.Id,
			&rent.ItemID,
			&rent.LesseeID,
			&rent.LessorID,
			&rent.StartDate,
			&rent.EndDate,
			&rent.TotalAmount,
			&rent.Status,
			&rent.PaymentStatus,
			&rent.DeliveryMethod,
			&rent.Notes,
			&rent.CreatedAt,
			&rent.UpdatedAt,
			&rent.StartedAt,
			&rent.CompletedAt,
			&rent.CancelledAt,
		)
		if err != nil {
			return nil, err
		}

		rentals = append(rentals, &rent)
	}

	if len(rentals) == 0 {
		return nil, rental.ErrNoRentalsFound
	}

	return rentals, nil
}

// ListByLessor implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) ListByLessor(ctx context.Context, lessorID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	var rentals []*rental.Rental

	query := `SELECT id, item_id, lessee_id, lessor_id, start_date, end_date, total_amount, status, payment_status, delivery_method, notes, created_at, updated_at, started_at, completed_at, cancelled_at
			  FROM rentals WHERE lessor_id = $1`

	args := []interface{}{lessorID}

	if status != nil {
		query += " AND status = $2"
		args = append(args, *status)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rent rental.Rental

		err := rows.Scan(
			&rent.Id,
			&rent.ItemID,
			&rent.LesseeID,
			&rent.LessorID,
			&rent.StartDate,
			&rent.EndDate,
			&rent.TotalAmount,
			&rent.Status,
			&rent.PaymentStatus,
			&rent.DeliveryMethod,
			&rent.Notes,
			&rent.CreatedAt,
			&rent.UpdatedAt,
			&rent.StartedAt,
			&rent.CompletedAt,
			&rent.CancelledAt,
		)
		if err != nil {
			return nil, err
		}

		rentals = append(rentals, &rent)
	}

	if len(rentals) == 0 {
		return nil, rental.ErrNoRentalsFound
	}

	return rentals, nil
}

func (r *RentalRepository) GetAllUserRentals(ctx context.Context, userID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	var rentals []*rental.Rental

	query := `SELECT id, item_id, lessee_id, lessor_id, start_date, end_date, total_amount, status, payment_status, delivery_method, notes, created_at, updated_at, started_at, completed_at, cancelled_at
			  FROM rentals WHERE (lessor_id = $1 OR lessee_id = $1)`

	args := []interface{}{userID}

	if status != nil {
		query += " AND status = $2"
		args = append(args, *status)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rent rental.Rental

		err := rows.Scan(
			&rent.Id,
			&rent.ItemID,
			&rent.LesseeID,
			&rent.LessorID,
			&rent.StartDate,
			&rent.EndDate,
			&rent.TotalAmount,
			&rent.Status,
			&rent.PaymentStatus,
			&rent.DeliveryMethod,
			&rent.Notes,
			&rent.CreatedAt,
			&rent.UpdatedAt,
			&rent.StartedAt,
			&rent.CompletedAt,
			&rent.CancelledAt,
		)
		if err != nil {
			return nil, err
		}

		rentals = append(rentals, &rent)
	}

	if len(rentals) == 0 {
		return nil, rental.ErrNoRentalsFound
	}

	return rentals, nil
}

// Update implements [rental.InterfaceRentalRepository].
func (r *RentalRepository) Update(ctx context.Context, rent *rental.Rental) error {
	query := `
		UPDATE rentals SET
			item_id = $1,
			lessee_id = $2,
			lessor_id = $3,
			start_date = $4,
			end_date = $5,
			total_amount = $6,
			status = $7,
			payment_status = $8,
			delivery_method = $9,
			notes = $10,
			update_at = $11,
			started_at = $12,
			complete_at = $13,
			cancelled_at = $14
		WHERE id = $15
	`

	result, err := r.db.ExecContext(ctx, query,
		rent.ItemID,
		rent.LesseeID,
		rent.LessorID,
		rent.StartDate,
		rent.EndDate,
		rent.TotalAmount,
		rent.Status,
		rent.PaymentStatus,
		rent.DeliveryMethod,
		rent.Notes,
		rent.UpdatedAt,
		rent.StartedAt,
		rent.CompletedAt,
		rent.CancelledAt,
		rent.Id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return rental.ErrRentalNotFound
	}

	return nil
}

func NewRentalRepository(db *sql.DB) rental.InterfaceRentalRepository {
	return &RentalRepository{db: db}
}
