package authaccount

import "errors"

var (
	ErrAuthAccountNotFound      = errors.New("auth account not found")
	ErrAuthAccountAlreadyExists = errors.New("auth account already exists")

	ErrEmailAlreadyInUse = errors.New("email already in use")

	ErrInvalidProvider = errors.New("invalid provider")

	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrPasswordRequired = errors.New("password is required for local provider")

	ErrEmailNotVerified = errors.New("email is not verified")
)
