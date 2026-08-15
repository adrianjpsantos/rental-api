package validator

import (
	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/go-playground/validator/v10"
)

func ValidateProvider(fl validator.FieldLevel) bool {
	provider := fl.Field().String()

	return IsProvider(provider)
}

func IsProvider(provider string) bool {
	switch provider {
	case string(authaccount.ProviderLocal),
		string(authaccount.ProviderGoogle),
		string(authaccount.ProviderFacebook),
		string(authaccount.ProviderGithub),
		string(authaccount.ProviderApple), string(authaccount.ProviderMicrosoft):
		return true
	default:
		return false

	}
}
