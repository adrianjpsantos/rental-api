package validator

import (
	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/go-playground/validator/v10"
)

func validateProvider(fl validator.FieldLevel) bool {
	role := fl.Field().String()

	switch role {
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
