package profile

import "errors"

var (
	ErrProfileNotFound      = errors.New("profile not found")
	ErrProfileAlreadyExists = errors.New("profile already exists")

	ErrInvalidName      = errors.New("invalid name")
	ErrInvalidCPF       = errors.New("invalid cpf")
	ErrInvalidPhone     = errors.New("invalid phone")
	ErrInvalidBirthDate = errors.New("invalid birth date")
	ErrInvalidAvatarURL = errors.New("invalid avatar url")
)
