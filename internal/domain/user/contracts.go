package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserForAuthentication(ctx context.Context, email string) (*UserForAuthentication, error)
	GetByEmail(ctx context.Context, email string) (*UserPublic, error)
	GetByCPF(ctx context.Context, cpf string) (*UserPublic, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error // soft delete
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByCPF(ctx context.Context, cpf string) (bool, error)
	UpdateReputationCache(ctx context.Context, id uuid.UUID) error
	UpdateTotalRentalCache(ctx context.Context, id uuid.UUID) error
	UpdateTotalItemsRentedCache(ctx context.Context, id uuid.UUID) error
}

type Service interface {
	Create(ctx context.Context, input UserCreateInput) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserForAuthentication(ctx context.Context, email string) (*UserForAuthentication, error)
	GetByEmail(ctx context.Context, email string) (*UserPublic, error)
	GetByCPF(ctx context.Context, cpf string) (*UserPublic, error)
	Update(ctx context.Context, id uuid.UUID, input UserUpdateInput) error
	Delete(ctx context.Context, id uuid.UUID) error // soft delete
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByCPF(ctx context.Context, cpf string) (bool, error)
	UpdateReputationCache(ctx context.Context, id uuid.UUID) error
	UpdateTotalRentalCache(ctx context.Context, id uuid.UUID) error
	UpdateTotalItemsRentedCache(ctx context.Context, id uuid.UUID) error
}
