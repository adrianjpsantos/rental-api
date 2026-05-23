package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/google/uuid"
)

type RentalService struct {
	repository rental.Repository
}

func (s *RentalService) ExistsOverlapping(ctx context.Context, itemID uuid.UUID, startDate time.Time, endDate time.Time) (bool, error) {
	return s.repository.ExistsOverlapping(ctx, itemID, startDate, endDate)
}

func NewRentalService(repo rental.Repository) rental.Service {
	return &RentalService{
		repository: repo,
	}
}

func (s *RentalService) Create(ctx context.Context, input rental.RentalCreateInput) error {

	exists, err := s.repository.ExistsOverlapping(ctx, input.ItemID, input.StartDate, input.EndDate)
	if err != nil {
		return err
	}
	if exists {
		return rental.ErrRentalAlreadyExists
	}

	newRental, err := rental.NewRental(input)
	if err != nil {
		return err
	}

	err = s.repository.Create(ctx, *newRental)

	if err != nil {
		return err
	}

	return nil
}

func (s *RentalService) GetByID(ctx context.Context, id uuid.UUID) (*rental.Rental, error) {
	existingRental, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingRental == nil {
		return nil, rental.ErrRentalNotFound
	}
	return existingRental, nil
}

func (s *RentalService) ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *rental.RentalStatus) ([]*rental.Rental, error) {
	listRental, err := s.repository.ListByLessee(ctx, lesseeID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) ListByLessor(ctx context.Context, lessorID uuid.UUID, status *rental.RentalStatus) ([]*rental.Rental, error) {
	listRental, err := s.repository.ListByLessor(ctx, lessorID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) GetAllUserRentals(ctx context.Context, lessorID uuid.UUID, status *rental.RentalStatus) ([]*rental.Rental, error) {
	listRental, err := s.repository.ListByLessor(ctx, lessorID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) UpdateStatus(ctx context.Context, id uuid.UUID, status *rental.RentalStatus) error {

	existingRental, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingRental == nil {
		return rental.ErrRentalNotFound
	}

	existingRental.UpdateStatus(*status)

	err = s.repository.Update(ctx, *existingRental)
	if err != nil {
		return err
	}

	return nil
}

func (s *RentalService) Cancel(ctx context.Context, id uuid.UUID, cancellationReason string) error {

	existingRental, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingRental == nil {
		return rental.ErrRentalNotFound
	}

	existingRental.UpdateStatus(rental.Cancelled)
	existingRental.CancellationReason = cancellationReason

	err = s.repository.Update(ctx, *existingRental)
	if err != nil {
		return err
	}

	return nil
}
