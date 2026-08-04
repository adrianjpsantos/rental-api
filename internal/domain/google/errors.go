package google

import "errors"

var (
	ErrInvalidIDToken = errors.New(
		"ID Token inválido",
	)

	ErrEmptyIDToken = errors.New(
		"ID Token não fornecido",
	)

	ErrEmailNotVerified = errors.New(
		"Email não verificado",
	)
)
