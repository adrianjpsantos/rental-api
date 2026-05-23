package rental

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, rental Rental) error
	GetByID(ctx context.Context, id uuid.UUID) (*Rental, error)
	ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	ListByLessor(ctx context.Context, lessorID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	GetAllUserRentals(ctx context.Context, userID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	Update(ctx context.Context, rental Rental) error
	ExistsOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}

type Service interface {
	Create(ctx context.Context, rental RentalCreateInput) error
	GetByID(ctx context.Context, id uuid.UUID) (*Rental, error)
	ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	ListByLessor(ctx context.Context, lessorID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	GetAllUserRentals(ctx context.Context, userID uuid.UUID, status *RentalStatus) ([]*Rental, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status *RentalStatus) error
	ExistsOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
	Cancel(ctx context.Context, id uuid.UUID, cancellationReason string) error
}
