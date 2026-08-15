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
	_ = v.RegisterValidation("cpf", ValidateCpf)
	_ = v.RegisterValidation("adult", ValidateAdult)
	_ = v.RegisterValidation("role", ValidateRole)
	_ = v.RegisterValidation("pass_strength", ValidatePasswordStrength)
	_ = v.RegisterValidation("provider", ValidateProvider)
}
