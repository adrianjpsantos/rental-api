package availability

import "errors"

var (
	// Erros de validação básica
	ErrInvalidItemID      = errors.New("item_id é obrigatório")
	ErrInvalidStartDate   = errors.New("data de início é obrigatória")
	ErrInvalidEndDate     = errors.New("data de término é obrigatória")
	ErrEndDateBeforeStart = errors.New("a data de término não pode ser anterior à data de início")
	ErrInvalidType        = errors.New("tipo de availability inválido (deve ser 'available' ou 'blocked')")
	ErrInvalidReason      = errors.New("motivo (reason) inválido")

	// Erros de negócio
	ErrSlotAlreadyExists      = errors.New("já existe um slot no período informado")
	ErrOverlappingWithBlocked = errors.New("existe um bloqueio no período solicitado")
	ErrCannotBlockPastDates   = errors.New("não é possível bloquear datas no passado")
	ErrSlotNotFound           = errors.New("slot de disponibilidade não encontrado")

	// Erros relacionados a aluguel
	ErrCannotCreateRentalSlot = errors.New("não foi possível criar slot de aluguel")
	ErrInvalidRentalPeriod    = errors.New("período de aluguel inválido para criação de slot")
)
