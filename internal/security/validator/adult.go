package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func validateAdult(fl validator.FieldLevel) bool {
	birthDate, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	today := time.Now()

	eighteenthBirthday := birthDate.AddDate(18, 0, 0)

	return !today.Before(eighteenthBirthday)
}
