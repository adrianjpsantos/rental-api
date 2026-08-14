// validator.go
package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
	setupCustomValidations(validate)
}

func Get() *validator.Validate {
	return validate
}

func setupCustomValidations(v *validator.Validate) {
	// Suas validações customizadas
	_ = v.RegisterValidation("cpf", validateCpf)
	_ = v.RegisterValidation("adult", validateAdult)
	_ = v.RegisterValidation("role", validateRole)
	_ = v.RegisterValidation("pass_strength", validatePasswordStrength)
	_ = v.RegisterValidation("provider", validateProvider)
}
