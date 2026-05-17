package middlewares

import "github.com/google/uuid"

type CurrentUser struct {
	ID    uuid.UUID
	Email string
}
