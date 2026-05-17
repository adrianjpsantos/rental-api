package user

import (
	"context"

	"github.com/google/uuid"
)

type InterfaceUserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetUserForAuthentication(ctx context.Context, email string) (*UserForAuthentication, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByCPF(ctx context.Context, cpf string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID uuid.UUID) error // soft delete
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByCPF(ctx context.Context, cpf string) (bool, error)
	UpdateReputationCache(ctx context.Context, userID uuid.UUID) error
	UpdateTotalRentalCache(ctx context.Context, userID uuid.UUID) error
	UpdateTotalItemsRentedCache(ctx context.Context, userID uuid.UUID) error
}
