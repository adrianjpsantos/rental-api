package validator

import (
	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/go-playground/validator/v10"
)

func validateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()

	switch role {
	case string(user.Lessee), string(user.Lessor), string(user.Admin):
		return true
	default:
		return false

	}
}
