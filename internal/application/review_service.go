package application

import (
	"context"
	"log"

	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

type ReviewService struct {
	reviewRepo  review.Repository
	userService user.Service
}

func NewReviewService(reviewRepo review.Repository, userService user.Service) review.Service {
	return &ReviewService{
		reviewRepo:  reviewRepo,
		userService: userService,
	}
}

func (s *ReviewService) ExistsByRentalAndReviewer(ctx context.Context, rentalID uuid.UUID, reviewerID uuid.UUID) (bool, error) {
	return s.reviewRepo.ExistsByRentalAndReviewer(ctx, rentalID, reviewerID)
}

func (s *ReviewService) Create(ctx context.Context, input review.ReviewCreateInput) error {

	exists, err := s.ExistsByRentalAndReviewer(ctx, input.RentalID, input.ReviewerID)
	if err != nil {
		return err
	}
	if exists {
		return review.ErrReviewAlreadyExists
	}

	newReview, err := review.NewReview(input)
	if err != nil {
		return err
	}

	if err := s.reviewRepo.Create(ctx, *newReview); err != nil {
		return err
	}

	if err := s.userService.UpdateReputationCache(ctx, input.ReviewedID); err != nil {
		log.Println(err)
	}

	return nil
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
