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
	RentalID   uuid.UUID
	ReviewerID uuid.UUID
	ReviewedID uuid.UUID
	ItemID     uuid.UUID
	Rating     int
	Comment    string
	ReviewType ReviewType
	CreatedAt  time.Time

	// Índice único para evitar múltiplas avaliações do mesmo rental
	UniqueRentalReview string // apenas referência
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
