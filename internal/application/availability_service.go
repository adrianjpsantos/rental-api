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

func (s *AvailabilityService) Register(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time, slotType availability.AvailabilityType, reason availability.Reason) (*availability.AvailabilitySlot, error) {

	exists, err := s.availabilitySlotRepo.HasBlockingSlot(ctx, itemID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, availability.ErrSlotAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newSlot, err := availability.NewAvailabilitySlot(itemID, startDate, endDate, slotType, reason)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.availabilitySlotRepo.Create(ctx, newSlot); err != nil {
		return nil, err
	}

	return newSlot, nil
}

func (s *AvailabilityService) GetByID(ctx context.Context, slotID uuid.UUID) (*availability.AvailabilitySlot, error) {
	existingAvailability, err := s.availabilitySlotRepo.GetByID(ctx, slotID)
	if err != nil {
		return nil, err
	}
	if existingAvailability == nil {
		return nil, availability.ErrSlotNotFound
	}
	return existingAvailability, nil
}

func (s *AvailabilityService) FindByItemID(ctx context.Context, itemID uuid.UUID) ([]*availability.AvailabilitySlot, error) {
	listAvailability, err := s.availabilitySlotRepo.FindByItemID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	return listAvailability, nil
}

func (s *AvailabilityService) FindOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*availability.AvailabilitySlot, error) {
	listAvailability, err := s.availabilitySlotRepo.FindOverlapping(ctx, itemID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return listAvailability, nil
}

func (s *AvailabilityService) Delete(ctx context.Context, slotID uuid.UUID) error {
	err := s.availabilitySlotRepo.Delete(ctx, slotID)

	if err != nil {
		return err
	}

	return nil
}
