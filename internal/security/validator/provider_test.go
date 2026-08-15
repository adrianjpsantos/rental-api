package validator_test

import (
	"testing"

	authaccount "github.com/adrianjpsantos/rental-api/internal/domain/auth_account"
	"github.com/adrianjpsantos/rental-api/internal/security/validator"
)

func TestProvider(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Deve retornar true para Provider Local",
			input:    string(authaccount.ProviderLocal),
			expected: true,
		},
		{
			name:     "Deve retornar true para Provider Google",
			input:    string(authaccount.ProviderGoogle),
			expected: true,
		},
		{
			name:     "Deve retornar true para Provider Facebook",
			input:    string(authaccount.ProviderFacebook),
			expected: true,
		},
		{
			name:     "Deve retornar true para Provider Github",
			input:    string(authaccount.ProviderGithub),
			expected: true,
		},
		{
			name:     "Deve retornar true para Provider Apple",
			input:    string(authaccount.ProviderApple),
			expected: true,
		},
		{
			name:     "Deve retornar true para Provider Microsoft",
			input:    string(authaccount.ProviderMicrosoft),
			expected: true,
		},
		{
			name:     "Deve retornar false para Provider inválido ou vazio",
			input:    "twitter",
			expected: false,
		},
		{
			name:     "Deve retornar false para string vazia",
			input:    "",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.IsProvider(tc.input)

			if result != tc.expected {
				t.Errorf("Para o input '%s', esperava %v, mas recebeu %v", tc.input, tc.expected, result)
			}
		})
	}
}
