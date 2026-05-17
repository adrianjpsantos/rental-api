package authenticate

import "errors"

var (
	ErrInvalidCredentials = errors.New(
		"Email ou senha inválidos",
	)

	ErrInvalidToken = errors.New(
		"token inválido",
	)

	ErrExpiredToken = errors.New(
		"token expirado",
	)

	ErrMalformedToken = errors.New(
		"token malformado",
	)

	ErrMissingAuthorizationHeader = errors.New(
		"cabeçalho de autorização ausente",
	)

	ErrInvalidAuthorizationHeader = errors.New(
		"cabeçalho de autorização inválido",
	)

	ErrInvalidRefreshToken = errors.New(
		"refresh token inválido",
	)

	ErrRefreshTokenExpired = errors.New(
		"refresh token expirado",
	)

	ErrUnauthorized = errors.New(
		"não autorizado",
	)

	ErrForbidden = errors.New(
		"acesso negado",
	)
)
