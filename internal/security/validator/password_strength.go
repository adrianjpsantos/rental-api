package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validatePasswordStrength(fl validator.FieldLevel) bool {
	strengthRegex := regexp.MustCompile(`/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&^])[A-Za-z\d@$!%*?&^]{8,}$/`)
	password := fl.Field().String()

	return strengthRegex.MatchString(password)
}
