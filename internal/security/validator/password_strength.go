package validator

import (
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var (
		hasLower   bool
		hasUpper   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true

		case unicode.IsUpper(char):
			hasUpper = true

		case unicode.IsDigit(char):
			hasNumber = true

		case strings.ContainsRune("@$!%*?&^", char):
			hasSpecial = true
		}
	}

	return hasLower &&
		hasUpper &&
		hasNumber &&
		hasSpecial
}
