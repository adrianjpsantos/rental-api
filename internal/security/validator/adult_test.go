package validator_test

import (
	"testing"

	"github.com/adrianjpsantos/rental-api/internal/security/validator"
)

func TestAdultAge(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{

		{
			name:     "Deve retornar falso para Data Vazia",
			input:    "",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Menor de 18",
			input:    "10/02/2020",
			expected: false,
		},
		{
			name:     "Deve retornar falso para Data Inválida",
			input:    "PP002255",
			expected: false,
		},
		{
			name:     "Deve retornar true para Maior de 18",
			input:    "10/06/1998",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := validator.IsAdult(tc.input)

			if result != tc.expected {
				t.Errorf("Para o input '%s', esperava %v, mas recebeu %v", tc.input, tc.expected, result)
			}
		})
	}
}
