package review

import (
	"context"

	"github.com/google/uuid"
)

type InterfaceReviewRepository interface {
	Create(ctx context.Context, review *Review) error
	GetByID(ctx context.Context, id uuid.UUID) (*Review, error)
	GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*Review, error)
	GetByReviewedID(ctx context.Context, reviewedID uuid.UUID, reviewType *ReviewType) ([]*Review, error)
	ExistsByRentalAndReviewer(ctx context.Context, rentalID, reviewerID uuid.UUID) (bool, error)
	ListUserReviews(ctx context.Context, userID uuid.UUID) ([]*Review, error)
}
