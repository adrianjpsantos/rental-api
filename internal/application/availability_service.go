package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/google/uuid"
)

type AvailabilityService struct {
	repository availability.Repository
}

func NewAvailabilityService(repository availability.Repository) availability.Service {
	return &AvailabilityService{
		repository: repository,
	}
}

func (s *AvailabilityService) Create(ctx context.Context, slot availability.AvailabilitySlotCreateInput) error {

	exists, err := s.repository.ExistsBlockingSlot(ctx, slot.ItemID, slot.StartDate, slot.EndDate)
	if err != nil {
		return err
	}
	if exists {
		return availability.ErrSlotAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newAvailabilitySlot, err := availability.NewAvailabilitySlot(slot)
	if err != nil {
		return err
	}

	// Persiste no banco
	if err := s.repository.Create(ctx, *newAvailabilitySlot); err != nil {
		return err
	}

	return nil
}

func (s *AvailabilityService) GetByID(ctx context.Context, id uuid.UUID) (*availability.AvailabilitySlot, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *AvailabilityService) ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*availability.AvailabilitySlot, error) {
	return s.repository.ListByItemID(ctx, itemID)
}

func (s *AvailabilityService) ListOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*availability.AvailabilitySlot, error) {
	return s.repository.ListOverlapping(ctx, itemID, startDate, endDate)
}

func (s *AvailabilityService) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *AvailabilityService) ExistsBlockingSlot(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error) {
	return s.repository.ExistsBlockingSlot(ctx, itemID, startDate, endDate)
}
