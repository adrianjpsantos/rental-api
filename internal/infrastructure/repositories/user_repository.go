package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	db DBTX
}

func (r *UserRepository) Create(ctx context.Context, role user.Role) (*uuid.UUID, error) {

	query := `
	INSERT INTO users 
	(role) 
	VALUES ($1) RETURNING id
	`
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, query,
		role,
	).Scan(&id)

	if err != nil {
		log.Printf("CREATE USER ERR: %s", err.Error())
		return nil, err
	}

	return &id, nil
}

// Delete implements [user.InterfaceUserRepository].
func (r *UserRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	query := `
	UPDATE users 
	(active=$2) 
	WHERE Id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		userID,
		false,
	)

	if err != nil {
		fmt.Printf("DELETE USER ERR: %s", err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	var u user.User

	query := `SELECT id,role,reputation,total_rentals,total_items_rented,active,created_at,updated_at,deleted_at FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&u.ID,
		&u.Role,
		&u.Reputation,
		&u.TotalRentals,
		&u.TotalItemsRented,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, user.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("❕USER BY ID: ", true)
	return &u, nil
}

// Update implements [user.InterfaceUserRepository].
func (r *UserRepository) Update(ctx context.Context, u user.User) error {
	query := `
	UPDATE users 
	(role=$2,reputation=$3,total_rentals=$4,total_items_rented=$5,created_at=$6,update_at=$7) 
	WHERE Id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		u.ID,
		u.Role,
		u.Reputation,
		u.TotalRentals,
		u.TotalItemsRented,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("CREATE USER ERR: %s", err.Error())
		return err
	}

	return nil
}

// UpdateTotalItemsRentedCache implements [user.InterfaceUserRepository].
func (r *UserRepository) UpdateTotalItemsRentedCache(ctx context.Context, userID uuid.UUID) error {
	query := `
UPDATE users 
SET total_items_rented = (
    SELECT COUNT(*)
    FROM rentals
    WHERE lessor_id = $1
)
WHERE id = $1
`
	_, err := r.db.ExecContext(ctx, query, userID)

	if err != nil {
		return err
	}

	return nil
}

// UpdateTotalRentalCache implements [user.InterfaceUserRepository].
func (r *UserRepository) UpdateTotalRentalCache(ctx context.Context, userID uuid.UUID) error {
	query := `
UPDATE users 
SET total_rentals = (
    SELECT COUNT(*)
    FROM rentals
    WHERE lessee_id = $1
)
WHERE id = $1
`

	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateReputationCache(ctx context.Context, userID uuid.UUID) error {

	query := `
	UPDATE users 
	SET reputation = (
    SELECT COALESCE(AVG(rating), 0)
    FROM review
    WHERE lessor_id = $1 OR lessee_id = $1 ) WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db DBTX) user.Repository {
	return &UserRepository{db: db}
}
