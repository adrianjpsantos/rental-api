package review

import "github.com/google/uuid"

type InterfaceReviewRepository interface {
	Create(review *Review) error
	GetByID(id uuid.UUID) (*Review, error)
	GetByRentalID(rentalID uuid.UUID) ([]*Review, error)
	GetByReviewedID(reviewedID uuid.UUID, reviewType *ReviewType) ([]*Review, error)
	ExistsByRentalAndReviewer(rentalID, reviewerID uuid.UUID) (bool, error)
	ListUserReviews(userID uuid.UUID) ([]*Review, error)
}
