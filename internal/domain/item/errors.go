package item

import "errors"

// =============================================
// Erros de Validação da Entidade Item
// =============================================

var (
	// Erros obrigatórios / campos inválidos
	ErrInvalidOwnerID       = errors.New("owner_id é obrigatório")
	ErrInvalidCategoryID    = errors.New("category_id é obrigatório")
	ErrInvalidTitle         = errors.New("título deve ter entre 5 e 100 caracteres")
	ErrInvalidDescription   = errors.New("descrição deve ter no mínimo 10 caracteres")
	ErrInvalidPricePerDay   = errors.New("preço por dia deve ser maior que zero")
	ErrInvalidPricePerHour  = errors.New("preço por hora deve ser maior que zero quando informado")
	ErrInvalidQuantity      = errors.New("quantidade deve ser maior que zero")
	ErrInvalidMinRentalDays = errors.New("quantidade mínima de dias deve ser maior ou igual a 1")
	ErrInvalidMaxRentalDays = errors.New("quantidade máxima de dias deve ser maior que a mínima")
	ErrInvalidCondition     = errors.New("condição inválida")

	// Erros de lógica de negócio
	ErrTitleTooShort          = errors.New("o título deve ter pelo menos 5 caracteres")
	ErrTitleTooLong           = errors.New("o título não pode exceder 100 caracteres")
	ErrDescriptionTooShort    = errors.New("a descrição deve ter pelo menos 10 caracteres")
	ErrPricePerDayTooLow      = errors.New("o preço por dia não pode ser menor que R$ 1,00")
	ErrQuantityZeroOrNegative = errors.New("a quantidade deve ser maior que zero")
	ErrMinRentalDaysInvalid   = errors.New("o mínimo de dias de aluguel deve ser pelo menos 1")
	ErrMaxRentalDaysInvalid   = errors.New("o máximo de dias deve ser maior ou igual ao mínimo")
	ErrMaxRentalDaysTooHigh   = errors.New("o máximo de dias de aluguel não pode exceder 365 dias")
	ErrInvalidLocation        = errors.New("localização é obrigatória")
	ErrNoPhotos               = errors.New("é obrigatório adicionar pelo menos uma foto do item")

	// Erros relacionados ao estado do Item
	ErrItemNotActive          = errors.New("este item não está ativo")
	ErrItemOutOfStock         = errors.New("item sem quantidade disponível")
	ErrItemAlreadyDeactivated = errors.New("o item já está desativado")
	ErrItemNotFound           = errors.New("item não encontrado")

	// Erros de permissão
	ErrNotOwnerOfItem         = errors.New("você não é o dono deste item")
	ErrCannotEditActiveRental = errors.New("não é possível editar um item que possui aluguel ativo")

	// Erros de consistência
	ErrInvalidDateRange = errors.New("o período mínimo não pode ser maior que o máximo")
	ErrInvalidBrand     = errors.New("marca inválida")
	ErrInvalidModel     = errors.New("modelo inválido")
)
