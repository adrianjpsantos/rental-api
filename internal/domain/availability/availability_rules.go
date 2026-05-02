package availability

import "time"

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
