package user

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID               uuid.UUID  `json:"id" validate:"required,uuid"`
	Role             Role       `json:"role" validate:"required,oneof=user admin"`
	Reputation       float32    `json:"reputation" validate:"gte=0,lte=5"`
	TotalRentals     int        `json:"total_rentals" validate:"gte=0"`
	TotalItemsRented int        `json:"total_items_rented" validate:"gte=0"`
	Active           bool       `json:"active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type CreateInput struct {
	Role Role `json:"role"`
}

func New(role Role) (*User, error) {
	user := &User{
		Role:             role,
		Active:           true,
		Reputation:       0,
		TotalRentals:     0,
		TotalItemsRented: 0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if user.Role == "" {
		user.Role = RoleUser
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Activate() {
	u.Active = true
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.Active = false
	u.UpdatedAt = time.Now()
}

func (u *User) AddReputation(rating float32) {
	if u.Reputation == 0 {
		u.Reputation = rating
	} else {
		u.Reputation = (u.Reputation + rating) / 2
	}

	u.UpdatedAt = time.Now()
}

func (u *User) IncrementRentals() {
	u.TotalRentals++
	u.UpdatedAt = time.Now()
}

func (u *User) IncrementItemsRented() {
	u.TotalItemsRented++
	u.UpdatedAt = time.Now()
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) Validate() error {
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
