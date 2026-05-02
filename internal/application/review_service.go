package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/google/uuid"
)

type ReviewService struct {
	reviewRepo review.InterfaceReviewRepository
}

func NewReviewService(reviewRepo review.InterfaceReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}

func (s *ReviewService) Register(ctx context.Context, rentalID uuid.UUID, reviewerID uuid.UUID, reviewedID uuid.UUID, itemID uuid.UUID, rating int, comment string, reviewType review.ReviewType) (*review.Review, error) {

	// Verifica se o email já existe
	exists, err := s.reviewRepo.ExistsByRentalAndReviewer(rentalID, reviewerID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, review.ErrReviewAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newReview, err := review.NewReview(rentalID, reviewerID, reviewedID, itemID, rating, comment, reviewType)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.reviewRepo.Create(newReview); err != nil {
		return nil, err
	}

	return newReview, nil
}

func (s *ReviewService) GetByID(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	existingReview, err := s.reviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existingReview == nil {
		return nil, review.ErrReviewNotFound
	}
	return existingReview, nil
}

func (s *ReviewService) GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*review.Review, error) {
	listReview, err := s.reviewRepo.GetByRentalID(rentalID)
	if err != nil {
		return nil, err
	}

	return listReview, nil
}

func (s *ReviewService) GetByReviewedID(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	listReview, err := s.reviewRepo.GetByReviewedID(reviewedID, reviewType)
	if err != nil {
		return nil, err
	}

	return listReview, nil
}

func (s *ReviewService) ListUserReviews(ctx context.Context, userID uuid.UUID) ([]*review.Review, error) {
	listReview, err := s.reviewRepo.ListUserReviews(userID)

	if err != nil {
		return nil, err
	}

	return listReview, nil
}
