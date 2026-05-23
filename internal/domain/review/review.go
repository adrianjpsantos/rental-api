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

type ReviewCreateInput struct {
	RentalID   uuid.UUID  `json:"rental_id"`
	ReviewerID uuid.UUID  `json:"reviewer_id"`
	ReviewedID uuid.UUID  `json:"reviewed_id"`
	ItemID     uuid.UUID  `json:"item_id"`
	Rating     int        `json:"rating"`
	Comment    string     `json:"comment"`
	ReviewType ReviewType `json:"review_type"`
}

func NewReview(input ReviewCreateInput) (*Review, error) {

	review := &Review{
		Id:         uuid.New(),
		RentalID:   input.RentalID,
		ReviewerID: input.ReviewerID,
		ReviewedID: input.ReviewedID,
		ItemID:     input.ItemID,
		Rating:     input.Rating,
		Comment:    input.Comment,
		ReviewType: input.ReviewType,
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

func ParseReviewType(value string) (*ReviewType, error) {
	switch ReviewType(value) {
	case AsLessor, AsLessee:
		rt := ReviewType(value)
		return &rt, nil
	default:
		return nil, ErrInvalidReviewType
	}
}

// CanBeReviewedBy verifica se o usuário pode avaliar o aluguel
func (r *Review) CanBeReviewedBy(userID uuid.UUID) bool {
	return r.ReviewerID == userID
}

// IsFromLessee retorna se a avaliação foi feita pelo locatário (sobre o locador)
func (r *Review) IsFromLessee() bool {
	return r.ReviewType == AsLessee
}

// IsFromLessor retorna se a avaliação foi feita pelo locador (sobre o locatário)
func (r *Review) IsFromLessor() bool {
	return r.ReviewType == AsLessor
}

// CanReviewRental verifica regras gerais para avaliar um aluguel
func (r *Review) CanReviewRental() error {
	if !r.IsValidRating() {
		return ErrInvalidRating
	}
	return nil
}
