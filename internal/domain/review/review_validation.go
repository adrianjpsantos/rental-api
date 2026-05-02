package review

import "github.com/google/uuid"

// Validate valida as regras de negócio da avaliação
func (r *Review) Validate() error {
	if r.RentalID == uuid.Nil {
		return ErrInvalidRentalID
	}
	if r.ReviewerID == uuid.Nil {
		return ErrInvalidReviewerID
	}
	if r.ReviewedID == uuid.Nil {
		return ErrInvalidReviewedID
	}
	if r.ItemID == uuid.Nil {
		return ErrInvalidItemID
	}
	if r.ReviewerID == r.ReviewedID {
		return ErrCannotReviewOwnRental
	}
	if !r.IsValidRating() {
		return ErrInvalidRating
	}
	if r.ReviewType != AsLessor && r.ReviewType != AsLessee {
		return ErrInvalidReviewType
	}

	return nil
}
