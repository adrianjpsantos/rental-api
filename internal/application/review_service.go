package application

import (
	"context"
	"log"

	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

type ReviewService struct {
	reviewRepo review.InterfaceReviewRepository
	userRepo   user.InterfaceUserRepository
}

func NewReviewService(reviewRepo review.InterfaceReviewRepository, userRepo user.InterfaceUserRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
		userRepo:   userRepo,
	}
}

func (s *ReviewService) Register(ctx context.Context, rentalID uuid.UUID, reviewerID uuid.UUID, reviewedID uuid.UUID, itemID uuid.UUID, rating int, comment string, reviewType review.ReviewType) (*review.Review, error) {

	// Verifica se o email já existe
	exists, err := s.reviewRepo.ExistsByRentalAndReviewer(ctx, rentalID, reviewerID)
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
	if err := s.reviewRepo.Create(ctx, newReview); err != nil {
		return nil, err
	}

	if err := s.userRepo.UpdateReputationCache(ctx, reviewedID); err != nil {
		log.Println(err)
	}

	return newReview, nil
}

func (s *ReviewService) GetByID(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	existingReview, err := s.reviewRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingReview == nil {
		return nil, review.ErrReviewNotFound
	}
	return existingReview, nil
}

func (s *ReviewService) GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*review.Review, error) {
	return s.reviewRepo.GetByRentalID(ctx, rentalID)
}

func (s *ReviewService) GetReceivedReviews(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.reviewRepo.GetReceivedReviews(ctx, reviewedID, reviewType)
}

func (s *ReviewService) GetGivenReviews(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.reviewRepo.GetGivenReviews(ctx, reviewedID, reviewType)
}

func (s *ReviewService) GetUserReviews(ctx context.Context, userID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.reviewRepo.GetUserReviews(ctx, userID, reviewType)
}
