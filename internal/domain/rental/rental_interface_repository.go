package rental

import (
	"time"

	"github.com/google/uuid"
)

type InterfaceRentalRepository interface {
	Create(rental *Rental) error
	GetByID(id uuid.UUID) (*Rental, error)
	ListByLessee(lesseeID uuid.UUID, status *Status) ([]*Rental, error)
	ListByLessor(lessorID uuid.UUID, status *Status) ([]*Rental, error)
	Update(rental *Rental) error
	UpdateStatus(id uuid.UUID, status Status) error
	ExistsOverlapping(itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}
