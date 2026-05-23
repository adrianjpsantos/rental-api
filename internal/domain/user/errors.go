package user

import "errors"

var (
	// Erros de validação básica
	ErrInvalidName         = errors.New("nome deve ter entre 3 e 100 caracteres")
	ErrInvalidEmail        = errors.New("email inválido")
	ErrInvalidCPF          = errors.New("CPF inválido")
	ErrInvalidPhone        = errors.New("telefone inválido")
	ErrInvalidBirthDate    = errors.New("data de nascimento inválida")
	ErrInvalidRole         = errors.New("role inválido")
	ErrInvalidPasswordHash = errors.New("hash da senha é obrigatório")

	// Erros de negócio
	ErrEmailAlreadyExists = errors.New("este email já está cadastrado")
	ErrCPFAlreadyExists   = errors.New("este CPF já está cadastrado")
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrUserDeleted        = errors.New("usuário está desativado")
	ErrCannotChangeRole   = errors.New("não é permitido alterar o role do usuário")

	// Erros de permissão
	ErrNotAuthorized = errors.New("usuário não autorizado")
	ErrAdminOnly     = errors.New("esta ação requer permissão de administrador")
)
