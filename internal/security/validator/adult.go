package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateAdult(fl validator.FieldLevel) bool {
	birthDate, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	return IsAdult(birthDate.String())
}

func IsAdult(birthDate string) bool {
	layout := "02/01/2006"
	today := time.Now()
	date, err := time.Parse(layout, birthDate)

	if err != nil {
		return false
	}

	eighteenthBirthday := date.AddDate(18, 0, 0)

	return !today.Before(eighteenthBirthday)
}
