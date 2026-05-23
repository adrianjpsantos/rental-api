package authenticate

import "errors"

var (
	ErrInvalidCredentials = errors.New(
		"Email ou senha inválidos",
	)

	ErrUnauthorized = errors.New(
		"não autorizado",
	)

	ErrForbidden = errors.New(
		"acesso negado",
	)
)
