package review

import "errors"

var (
	// Erros de validação básica
	ErrInvalidRentalID   = errors.New("rental_id é obrigatório")
	ErrInvalidReviewerID = errors.New("reviewer_id é obrigatório")
	ErrInvalidReviewedID = errors.New("reviewed_id é obrigatório")
	ErrInvalidItemID     = errors.New("item_id é obrigatório")
	ErrInvalidRating     = errors.New("a avaliação deve ser entre 1 e 5 estrelas")
	ErrInvalidReviewType = errors.New("tipo de avaliação inválido")

	// Erros de negócio
	ErrReviewAlreadyExists      = errors.New("você já avaliou este aluguel")
	ErrCannotReviewFutureRental = errors.New("não é possível avaliar um aluguel que ainda não foi concluído")
	ErrCannotReviewOwnRental    = errors.New("não é possível avaliar a si mesmo")
	ErrRentalNotCompleted       = errors.New("apenas aluguéis concluídos podem ser avaliados")
	ErrReviewNotFound           = errors.New("avaliação não encontrada")

	// Erros de permissão
	ErrNotAuthorizedToReview     = errors.New("você não tem permissão para avaliar este aluguel")
	ErrReviewerMustBeParticipant = errors.New("o avaliador deve ser participante do aluguel")
)
