package rental

import "errors"

var (
	// Erros de validação básica
	ErrInvalidItemID         = errors.New("item_id é obrigatório")
	ErrInvalidLesseeID       = errors.New("lessee_id é obrigatório")
	ErrInvalidLessorID       = errors.New("lessor_id é obrigatório")
	ErrInvalidStartDate      = errors.New("data de início é obrigatória")
	ErrInvalidEndDate        = errors.New("data de término é obrigatória")
	ErrInvalidTotalAmount    = errors.New("valor total deve ser maior que zero")
	ErrInvalidDeliveryMethod = errors.New("método de entrega inválido")

	// Erros de lógica de datas
	ErrStartDateInPast      = errors.New("a data de início não pode ser no passado")
	ErrEndDateBeforeStart   = errors.New("a data de término não pode ser anterior à data de início")
	ErrRentalPeriodTooShort = errors.New("o período mínimo de aluguel não foi respeitado")
	ErrRentalPeriodTooLong  = errors.New("o período de aluguel excede o máximo permitido")

	// Erros de status
	ErrInvalidStatusTransition = errors.New("transição de status inválida")
	ErrCannotCancelRental      = errors.New("não é possível cancelar este aluguel no status atual")

	// Erros de negócio
	ErrRentalNotFound      = errors.New("contrato de aluguel não encontrado")
	ErrItemNotAvailable    = errors.New("item não está disponível no período solicitado")
	ErrLesseeIsOwner       = errors.New("o locatário não pode ser o próprio dono do item")
	ErrRentalAlreadyExists = errors.New("já existe um aluguel ativo para este item no período")

	// Erros de permissão
	ErrNotAuthorizedToApprove = errors.New("apenas o locador pode aprovar o aluguel")
	ErrNotAuthorizedToCancel  = errors.New("você não tem permissão para cancelar este aluguel")

	// === Erros de Listagem de Contratos ===
	ErrLesseeHasNoRentals = errors.New("este locatário ainda não possui nenhum contrato de aluguel")
	ErrLessorHasNoRentals = errors.New("este locador ainda não possui nenhum item alugado")

	ErrNoRentalsFound          = errors.New("nenhum contrato de aluguel encontrado")
	ErrNoActiveRentalsFound    = errors.New("nenhum contrato de aluguel ativo encontrado")
	ErrNoPendingRentalsFound   = errors.New("nenhum contrato de aluguel pendente encontrado")
	ErrNoCompletedRentalsFound = errors.New("nenhum contrato de aluguel concluído encontrado")

	// Erros de contexto
	ErrUserHasNoRentalsAsLessee = errors.New("o usuário não tem histórico como locatário")
	ErrUserHasNoRentalsAsLessor = errors.New("o usuário não tem histórico como locador")
)
