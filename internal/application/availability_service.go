package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/availability"
	"github.com/google/uuid"
)

type AvailabilityService struct {
	availabilitySlotRepo availability.InterfaceAvailabilityRepository
}

func NewAvailabilityService(availabilitySlotRepo availability.InterfaceAvailabilityRepository) *AvailabilityService {
	return &AvailabilityService{
		availabilitySlotRepo: availabilitySlotRepo,
	}
}

func (s *AvailabilityService) Register(ctx context.Context, newSlot availability.AvailabilitySlotCreateInput) (*availability.AvailabilitySlot, error) {

	exists, err := s.availabilitySlotRepo.ExistsBlockingSlot(ctx, newSlot.ItemID, newSlot.StartDate, newSlot.EndDate)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, availability.ErrSlotAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newAvailabilitySlot, err := availability.NewAvailabilitySlot(newSlot)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.availabilitySlotRepo.Create(ctx, newAvailabilitySlot); err != nil {
		return nil, err
	}

	return newAvailabilitySlot, nil
}

func (s *AvailabilityService) GetByID(ctx context.Context, slotID uuid.UUID) (*availability.AvailabilitySlot, error) {
	return s.availabilitySlotRepo.GetByID(ctx, slotID)
}

func (s *AvailabilityService) ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*availability.AvailabilitySlot, error) {
	return s.availabilitySlotRepo.ListByItemID(ctx, itemID)
}

func (s *AvailabilityService) ListOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*availability.AvailabilitySlot, error) {
	return s.availabilitySlotRepo.ListOverlapping(ctx, itemID, startDate, endDate)
}

func (s *AvailabilityService) Delete(ctx context.Context, slotID uuid.UUID) error {
	err := s.availabilitySlotRepo.Delete(ctx, slotID)

	if err != nil {
		return err
	}

	return nil
}

func (s *AvailabilityService) ExistBlockingSlot(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) (bool, error) {
	return s.availabilitySlotRepo.ExistsBlockingSlot(ctx, itemID, startDate, endDate)
}
