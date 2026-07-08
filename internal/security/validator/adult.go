package validator

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

func validateAdult(fl validator.FieldLevel) bool {
	const layout = "02/01/2006"

	birthDate, err := time.Parse(layout, fl.Field().String())
	if err != nil {
		fmt.Printf("Error Validate BirthDate: %s\n", err.Error())
		return false
	}

	now := time.Now()
	eighteenYearsAgo := now.AddDate(-18, 0, 0)

	return !birthDate.After(eighteenYearsAgo)
}
