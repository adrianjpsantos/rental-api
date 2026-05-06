package availability

import (
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/google/uuid"
)

type AvailabilityType string

const (
	Blocked   AvailabilityType = "blocked"
	Available AvailabilityType = "available"
)

type Reason string

const (
	Maintenance Reason = "maintenance"
	Rented      Reason = "rented"
	OwnerBlock  Reason = "owner_block"
)

type AvailabilitySlot struct {
	Id        uuid.UUID
	ItemId    uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	Type      AvailabilityType
	Reason    Reason
	CreatedAt time.Time

	// Índice composto (útil para consultas de disponibilidade)
	ItemIdStartEndIdx string // ignorado

	// Relacionamento
	Item item.Item
}

func NewAvailabilitySlot(
	itemID uuid.UUID,
	startDate, endDate time.Time,
	slotType AvailabilityType,
	reason Reason,
) (*AvailabilitySlot, error) {

	slot := &AvailabilitySlot{
		Id:        uuid.New(),
		ItemId:    itemID,
		StartDate: startDate,
		EndDate:   endDate,
		Type:      slotType,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	if err := slot.Validate(); err != nil {
		return nil, err
	}

	return slot, nil
}
