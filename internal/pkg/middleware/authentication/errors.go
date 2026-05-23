package authentication

import "errors"

var (
	ErrMissingAuthorizationHeader = errors.New("missing authorization header")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrInvalidSessionToken        = errors.New("invalid session token")
	ErrAuthenticatedUserNotFound  = errors.New("authenticated user not found")
)
