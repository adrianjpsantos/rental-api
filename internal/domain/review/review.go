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
	Id         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RentalID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	ReviewerID uuid.UUID  `gorm:"type:uuid;not null"`
	ReviewedID uuid.UUID  `gorm:"type:uuid;not null"`
	ItemID     uuid.UUID  `gorm:"type:uuid;not null"`
	Rating     int        `gorm:"not null;check:rating BETWEEN 1 AND 5"`
	Comment    string     `gorm:"type:text"`
	ReviewType ReviewType `gorm:"size:20;not null"`
	CreatedAt  time.Time

	// Índice único para evitar múltiplas avaliações do mesmo rental
	UniqueRentalReview string `gorm:"-"` // apenas referência
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
