package validator_test

import (
	"testing"

	"github.com/adrianjpsantos/rental-api/internal/security/validator"
)

func TestPasswordStrength(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Deve retornar falso para senha vazia",
			input:    "",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Senha com tamanho incorreto",
			input:    "123456",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Senha sem Letra Maiuscula",
			input:    "senha123",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Senha sem Letra Minuscula",
			input:    "SENHA123",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Senha sem Numeros",
			input:    "senhaAAA",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Senha sem Caracteres Especiais",
			input:    "Senha123",
			expected: false,
		},
		{
			name:     "Deve retornar true para Senha Forte",
			input:    "Senha.123",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.IsPasswordStrength(tc.input)

			if result != tc.expected {
				t.Errorf("Para o input '%s', esperava %v, mas recebeu %v", tc.input, tc.expected, result)
			}
		})
	}
}
