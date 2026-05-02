package availability

import (
	"time"

	"github.com/google/uuid"
)

type InterfaceAvailabilityRepository interface {
	Create(slot *AvailabilitySlot) error
	GetByID(id uuid.UUID) (*AvailabilitySlot, error)
	FindByItemID(itemID uuid.UUID) ([]*AvailabilitySlot, error)
	FindOverlapping(itemID uuid.UUID, startDate, endDate time.Time) ([]*AvailabilitySlot, error)
	Delete(id uuid.UUID) error
	// Método útil para verificar disponibilidade rapidamente
	HasBlockingSlot(itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}
