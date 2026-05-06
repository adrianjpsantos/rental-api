package availability

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InterfaceAvailabilityRepository interface {
	Create(ctx context.Context, slot *AvailabilitySlot) error
	GetByID(ctx context.Context, id uuid.UUID) (*AvailabilitySlot, error)
	FindByItemID(ctx context.Context, itemID uuid.UUID) ([]*AvailabilitySlot, error)
	FindOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*AvailabilitySlot, error)
	Delete(ctx context.Context, id uuid.UUID) error
	// Método útil para verificar disponibilidade rapidamente
	HasBlockingSlot(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error)
}
