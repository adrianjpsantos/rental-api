package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/uow"
	"github.com/google/uuid"
)

type ReviewService struct {
	Repository  review.Repository
	userService user.Service
	UOW         uow.UnitOfWork
}

func NewReviewService(reviewRepo review.Repository, userService user.Service) review.Service {
	return &ReviewService{
		Repository:  reviewRepo,
		userService: userService,
	}
}

func (s *ReviewService) ExistsByRentalAndReviewer(ctx context.Context, rentalID uuid.UUID, reviewerID uuid.UUID) (bool, error) {
	return s.Repository.ExistsByRentalAndReviewer(ctx, rentalID, reviewerID)
}

func (s *ReviewService) Create(ctx context.Context, input review.ReviewCreateInput) error {

	exists, err := s.Repository.ExistsByRentalAndReviewer(ctx, input.RentalID, input.ReviewerID)
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

	if err := s.Repository.Create(ctx, *newReview); err != nil {
		return err
	}

	// Update the reputation of the reviewed user with Trigger in PostgreSQL, but also update the reputation in the user service to keep it in sync

	return nil
}

func (s *ReviewService) GetByID(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	existingReview, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingReview == nil {
		return nil, review.ErrReviewNotFound
	}
	return existingReview, nil
}

func (s *ReviewService) GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*review.Review, error) {
	return s.Repository.GetByRentalID(ctx, rentalID)
}

func (s *ReviewService) GetReceivedReviews(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.Repository.GetReceivedReviews(ctx, reviewedID, reviewType)
}

func (s *ReviewService) GetGivenReviews(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.Repository.GetGivenReviews(ctx, reviewedID, reviewType)
}

func (s *ReviewService) GetUserReviews(ctx context.Context, userID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	return s.Repository.GetUserReviews(ctx, userID, reviewType)
}
