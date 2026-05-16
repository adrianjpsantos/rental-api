package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/rental"
	"github.com/google/uuid"
)

type RentalService struct {
	rentalRepo rental.InterfaceRentalRepository
}

func NewRentalService(rentalRepo rental.InterfaceRentalRepository) *RentalService {
	return &RentalService{
		rentalRepo: rentalRepo,
	}
}

func (s *RentalService) Register(ctx context.Context, itemID, lesseeID, lessorID uuid.UUID, startDate, endDate time.Time, totalAmount float64, deliveryMethod rental.DeliveryMethod, notes string) (*rental.Rental, error) {

	exists, err := s.rentalRepo.ExistsOverlapping(ctx, itemID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, rental.ErrRentalAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newRental, err := rental.NewRental(itemID, lesseeID, lessorID, startDate, endDate, totalAmount, deliveryMethod, notes)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.rentalRepo.Create(ctx, newRental); err != nil {
		return nil, err
	}

	return newRental, nil
}

func (s *RentalService) GetByID(ctx context.Context, id uuid.UUID) (*rental.Rental, error) {
	existingRental, err := s.rentalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingRental == nil {
		return nil, rental.ErrRentalNotFound
	}
	return existingRental, nil
}

func (s *RentalService) ListByLessee(ctx context.Context, lesseeID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	listRental, err := s.rentalRepo.ListByLessee(ctx, lesseeID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) ListByLessor(ctx context.Context, lessorID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	listRental, err := s.rentalRepo.ListByLessor(ctx, lessorID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) GetAllUserRentals(ctx context.Context, lessorID uuid.UUID, status *rental.Status) ([]*rental.Rental, error) {
	listRental, err := s.rentalRepo.ListByLessor(ctx, lessorID, status)
	if err != nil {
		return nil, err
	}

	return listRental, nil
}

func (s *RentalService) UpdateStatus(ctx context.Context, userID uuid.UUID, newStatus *rental.Status) (*rental.Rental, error) {

	existingRental, err := s.rentalRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existingRental == nil {
		return nil, rental.ErrRentalNotFound
	}

	existingRental.UpdateStatus(*newStatus)

	// Persiste as alterações
	if err := s.rentalRepo.Update(ctx, existingRental); err != nil {
		return nil, err
	}

	return existingRental, nil
}

func (s *RentalService) Cancel(ctx context.Context, userID uuid.UUID, cancellationReason string) (*rental.Rental, error) {

	existingRental, err := s.rentalRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existingRental == nil {
		return nil, rental.ErrRentalNotFound
	}

	existingRental.UpdateStatus(rental.Cancelled)
	existingRental.CancellationReason = cancellationReason

	// Persiste as alterações
	if err := s.rentalRepo.Update(ctx, existingRental); err != nil {
		return nil, err
	}

	return existingRental, nil
}
