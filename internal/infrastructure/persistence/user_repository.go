package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {

	existsEmail, err := r.ExistsByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	if existsEmail {
		return user.ErrEmailAlreadyExists
	}

	existsCPF, err := r.ExistsByCPF(ctx, u.CPF)
	if err != nil {
		return err
	}
	if existsCPF {
		return user.ErrCPFAlreadyExists
	}

	///NOT EXIST

	query := `
	INSERT INTO users 
	(id,name,email,password_hash,cpf,phone,birth_date,avatar_url,is_verified,role,reputation,total_rentals,total_items_rented,created_at,update_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
	`

	_, err = r.db.ExecContext(ctx, query,
		u.Id,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.CPF,
		u.Phone,
		u.BirthDate,
		u.AvatarURL,
		u.IsVerified,
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

// Delete implements [user.InterfaceUserRepository].
func (r *UserRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	panic("unimplemented")
}

// ExistsByCPF implements [user.InterfaceUserRepository].
func (r *UserRepository) ExistsByCPF(ctx context.Context, cpf string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE cpf = $1)`
	err := r.db.QueryRowContext(ctx, query, cpf).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, user.ErrUserNotFound
	}

	if err != nil {
		return false, err
	}

	fmt.Println("❕USER EXISTS BY CPF: ", exists)
	return exists, nil
}

// ExistsByEmail implements [user.InterfaceUserRepository].
func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, user.ErrUserNotFound
	}

	if err != nil {
		return false, err
	}

	fmt.Println("❕USER EXISTS BY EMAIL: ", exists)
	return exists, nil
}

// GetByCPF implements [user.InterfaceUserRepository].
func (r *UserRepository) GetByCPF(ctx context.Context, cpf string) (*user.User, error) {
	var u user.User

	query := `SELECT id,name,email,cpf,phone,birth_date,avatar_url,is_verified,role,reputation,total_rentals,total_items_rented,created_at FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, cpf).Scan(&u.Id,
		&u.Name,
		&u.Email,
		&u.CPF,
		&u.Phone,
		&u.BirthDate,
		&u.AvatarURL,
		&u.IsVerified,
		&u.Role,
		&u.Reputation,
		&u.TotalRentals,
		&u.TotalItemsRented,
		&u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, user.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	fmt.Println("❕USER BY CPF: ", true)
	return &u, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User

	query := `SELECT id,name,email,cpf,phone,birth_date,avatar_url,is_verified,role,reputation,total_rentals,total_items_rented,created_at FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.Id,
		&u.Name,
		&u.Email,
		&u.CPF,
		&u.Phone,
		&u.BirthDate,
		&u.AvatarURL,
		&u.IsVerified,
		&u.Role,
		&u.Reputation,
		&u.TotalRentals,
		&u.TotalItemsRented,
		&u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, user.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("❕USER BY EMAIL: ", true)
	return &u, nil
}

func (r *UserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	var u user.User

	query := `SELECT id,name,email,cpf,phone,birth_date,avatar_url,is_verified,role,reputation,total_rentals,total_items_rented,created_at FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&u.Id,
		&u.Name,
		&u.Email,
		&u.CPF,
		&u.Phone,
		&u.BirthDate,
		&u.AvatarURL,
		&u.IsVerified,
		&u.Role,
		&u.Reputation,
		&u.TotalRentals,
		&u.TotalItemsRented,
		&u.CreatedAt)

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
func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	query := `
	UPDATE users 
	(id = $1,name=$2,email=$3,password_hash=$4,cpf=$5,phone=$6,birth_date=$7,avatar_url=$8,is_verified=$9,role=$10,reputation=$11,total_rentals=$12,total_items_rented=$13,created_at=$14,update_at=$15) 
	WHERE Id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		u.Id,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.CPF,
		u.Phone,
		u.BirthDate,
		u.AvatarURL,
		u.IsVerified,
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

func NewUserRepository(db *sql.DB) user.InterfaceUserRepository {
	return &UserRepository{db: db}
}
