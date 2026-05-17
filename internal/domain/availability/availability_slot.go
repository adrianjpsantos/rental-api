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

type AvailabilityFilter struct {
	ItemID             uuid.UUID
	StartDate, EndDate time.Time
	Type               *AvailabilityType
	Reason             *Reason
}

type AvailabilitySlot struct {
	Id                 uuid.UUID
	ItemId             uuid.UUID
	StartDate, EndDate time.Time
	Type               AvailabilityType
	Reason             Reason
	CreatedAt          time.Time

	// Índice composto (útil para consultas de disponibilidade)
	ItemIdStartEndIdx string // ignorado

	// Relacionamento
	Item item.Item
}

type AvailabilitySlotCreateInput struct {
	ItemID             uuid.UUID
	StartDate, EndDate time.Time
	SlotType           AvailabilityType
	Reason             Reason
}

func NewAvailabilitySlot(
	newSlot AvailabilitySlotCreateInput,
) (*AvailabilitySlot, error) {

	slot := &AvailabilitySlot{
		Id:        uuid.New(),
		ItemId:    newSlot.ItemID,
		StartDate: newSlot.StartDate,
		EndDate:   newSlot.EndDate,
		Type:      newSlot.SlotType,
		Reason:    newSlot.Reason,
		CreatedAt: time.Now(),
	}

	if err := slot.Validate(); err != nil {
		return nil, err
	}

	return slot, nil
}
