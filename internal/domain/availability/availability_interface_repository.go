package availability

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InterfaceAvailabilityRepository interface {
	Create(ctx context.Context, slot *AvailabilitySlot) error
	GetByID(ctx context.Context, id uuid.UUID) (*AvailabilitySlot, error)
	ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*AvailabilitySlot, error)
	ListOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*AvailabilitySlot, error)
	Delete(ctx context.Context, id uuid.UUID) error
	// Método útil para verificar disponibilidade rapidamente
	ExistsBlockingSlot(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}
