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

func (s *AvailabilitySlot) Validate() error {
	if s.ItemId == uuid.Nil {
		return ErrInvalidItemID
	}
	if s.StartDate.IsZero() {
		return ErrInvalidStartDate
	}
	if s.EndDate.IsZero() {
		return ErrInvalidEndDate
	}
	if s.EndDate.Before(s.StartDate) {
		return ErrEndDateBeforeStart
	}
	if s.Type != Available && s.Type != Blocked {
		return ErrInvalidType
	}
	if s.Reason == "" {
		return ErrInvalidReason
	}

	// Não permitir bloquear datas muito antigas
	if s.IsBlocked() && s.EndDate.Before(time.Now().AddDate(0, -1, 0)) { // mais de 1 mês atrás
		return ErrCannotBlockPastDates
	}

	return nil
}

// Overlaps verifica se dois slots se sobrepõem
func (s *AvailabilitySlot) Overlaps(other *AvailabilitySlot) bool {
	return s.StartDate.Before(other.EndDate) && s.EndDate.After(other.StartDate)
}

// Contains verifica se uma data está dentro do slot
func (s *AvailabilitySlot) Contains(date time.Time) bool {
	return (date.Equal(s.StartDate) || date.After(s.StartDate)) &&
		(date.Equal(s.EndDate) || date.Before(s.EndDate))
}

// IsBlocked retorna se o slot é de bloqueio
func (s *AvailabilitySlot) IsBlocked() bool {
	return s.Type == Blocked
}

// IsAvailable retorna se o slot é de disponibilidade
func (s *AvailabilitySlot) IsAvailable() bool {
	return s.Type == Available
}

// IsInPast verifica se o slot está totalmente no passado
func (s *AvailabilitySlot) IsInPast() bool {
	return s.EndDate.Before(time.Now())
}
