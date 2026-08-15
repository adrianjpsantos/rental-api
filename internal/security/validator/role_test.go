package validator_test

import (
	"testing"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/security/validator"
)

func TestRole(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Deve retornar true para Role Admin",
			input:    string(user.RoleAdmin),
			expected: true,
		},
		{
			name:     "Deve retornar true para Role User",
			input:    string(user.RoleUser),
			expected: true,
		},
		{
			name:     "Deve retornar false para Role desconhecida",
			input:    "super_admin",
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
			result := validator.IsRole(tc.input)

			if result != tc.expected {
				t.Errorf("Para o input '%s', esperava %v, mas recebeu %v", tc.input, tc.expected, result)
			}
		})
	}
}
