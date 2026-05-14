package review

import (
	"context"

	"github.com/google/uuid"
)

type InterfaceReviewRepository interface {
	Create(ctx context.Context, review *Review) error
	GetByID(ctx context.Context, id uuid.UUID) (*Review, error)
	GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*Review, error)
	GetReceivedReviews(ctx context.Context, reviewedID uuid.UUID, reviewType *ReviewType) ([]*Review, error)
	GetGivenReviews(ctx context.Context, reviewerID uuid.UUID, reviewType *ReviewType) ([]*Review, error)
	ExistsByRentalAndReviewer(ctx context.Context, rentalID, reviewerID uuid.UUID) (bool, error)
	GetUserReviews(ctx context.Context, userID uuid.UUID, reviewType *ReviewType) ([]*Review, error)
}
