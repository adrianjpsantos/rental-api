package availability

import (
	"time"

	"github.com/google/uuid"
)

// Validate valida as regras de negócio do AvailabilitySlot
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
