package validator

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/go-playground/validator/v10"
)

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	return IsRole(role)
}

func IsRole(role string) bool {
	switch role {
	case string(user.RoleAdmin), string(user.RoleUser):
		return true
	default:
		return false

	}
}
