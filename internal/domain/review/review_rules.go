package review

import "github.com/google/uuid"

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
