package rental

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InterfaceRentalRepository interface {
	Create(ctx context.Context, rental *Rental) error
	GetByID(ctx context.Context, id uuid.UUID) (*Rental, error)
	ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *Status) ([]*Rental, error)
	ListByLessor(ctx context.Context, lessorID uuid.UUID, status *Status) ([]*Rental, error)
	GetAllUserRentals(ctx context.Context, userID uuid.UUID, status *Status) ([]*Rental, error)
	Update(ctx context.Context, rental *Rental) error
	ExistsOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}
