package review

import (
	"time"

	"github.com/google/uuid"
)

type ReviewType string

const (
	AsLessor ReviewType = "as_lessor" // Avaliação feita pelo locatário sobre o locador
	AsLessee ReviewType = "as_lessee" // Avaliação feita pelo locador sobre o locatário
)

type Review struct {
	Id         uuid.UUID
	RentalID   uuid.UUID  `json:"rental_id"`
	ReviewerID uuid.UUID  `json:"reviewer_id"`
	ReviewedID uuid.UUID  `json:"reviewed_id"`
	ItemID     uuid.UUID  `json:"item_id"`
	Rating     int        `json:"rating"`
	Comment    string     `json:"comment"`
	ReviewType ReviewType `json:"review_type"`
	CreatedAt  time.Time  `json:"created_at"`

	// Índice único para evitar múltiplas avaliações do mesmo rental
	UniqueRentalReview string `json:"unique_rental_review"` // apenas referência
}

func NewReview(
	rentalID uuid.UUID,
	reviewerID uuid.UUID,
	reviewedID uuid.UUID,
	itemID uuid.UUID,
	rating int,
	comment string,
	reviewType ReviewType,
) (*Review, error) {

	review := &Review{
		Id:         uuid.New(),
		RentalID:   rentalID,
		ReviewerID: reviewerID,
		ReviewedID: reviewedID,
		ItemID:     itemID,
		Rating:     rating,
		Comment:    comment,
		ReviewType: reviewType,
		CreatedAt:  time.Now(),
	}

	if err := review.Validate(); err != nil {
		return nil, err
	}

	return review, nil
}

// IsValidRating verifica se a nota está entre 1 e 5
func (r *Review) IsValidRating() bool {
	return r.Rating >= 1 && r.Rating <= 5
}
