package session

import "errors"

var (
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
)
