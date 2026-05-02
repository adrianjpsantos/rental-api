package persistence

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type rentalRepository struct {
	db *gorm.DB
}

// Create implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) Create(rental *rental.Rental) error {
	panic("unimplemented")
}

// ExistsOverlapping implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) ExistsOverlapping(itemID uuid.UUID, startDate time.Time, endDate time.Time) (bool, error) {
	panic("unimplemented")
}

// GetByID implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) GetByID(id uuid.UUID) (*rental.Rental, error) {
	panic("unimplemented")
}

// ListByLessee implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) ListByLessee(lesseeID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	panic("unimplemented")
}

// ListByLessor implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) ListByLessor(lessorID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	panic("unimplemented")
}

// Update implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) Update(rental *rental.Rental) error {
	panic("unimplemented")
}

// UpdateStatus implements [rental.InterfaceRentalRepository].
func (r *rentalRepository) UpdateStatus(id uuid.UUID, status rental.Status) error {
	panic("unimplemented")
}

func NewRentalRepository(db *gorm.DB) rental.InterfaceRentalRepository {
	return &rentalRepository{db: db}
}
