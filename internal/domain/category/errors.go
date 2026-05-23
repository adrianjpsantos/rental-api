package category

import "errors"

var (
	// Erros de validação
	ErrInvalidName        = errors.New("nome da categoria é obrigatório")
	ErrNameTooShort       = errors.New("nome da categoria deve ter pelo menos 3 caracteres")
	ErrNameTooLong        = errors.New("nome da categoria não pode ter mais que 80 caracteres")
	ErrInvalidSlug        = errors.New("slug da categoria é obrigatório")
	ErrSlugInvalidFormat  = errors.New("slug deve conter apenas letras minúsculas, números e hífens")
	ErrDescriptionTooLong = errors.New("descrição não pode ter mais que 500 caracteres")

	// Erros de negócio
	ErrCategoryNotFound    = errors.New("categoria não encontrada")
	ErrNameAlreadyExists   = errors.New("já existe uma categoria com este nome")
	ErrSlugAlreadyExists   = errors.New("já existe uma categoria com este slug")
	ErrCategoryHasItems    = errors.New("não é possível excluir uma categoria que possui itens cadastrados")
	ErrCannotDeleteDefault = errors.New("não é possível excluir uma categoria padrão do sistema")
	ErrCategoryInactive    = errors.New("categoria está inativa")
)
