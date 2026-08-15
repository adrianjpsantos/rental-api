package validator_test

import (
	"testing"

	"github.com/adrianjpsantos/rental-api/internal/security/validator"
)

// O nome da função de teste SEMPRE começa com a palavra Test seguida pelo nome da função
func TestIsCpf(t *testing.T) {
	// Usamos "Table-Driven Tests" (Testes guiados por tabela), que é o padrão de ouro em Go.
	// Criamos uma lista com vários cenários (casos de sucesso e de erro).
	tests := []struct {
		name     string // Descrição do teste
		input    string // O CPF que vamos enviar
		expected bool   // O resultado que esperamos (true ou false)
	}{
		{
			name:     "Deve retornar falso para CPF vazio",
			input:    "",
			expected: false,
		},
		{
			name:     "Deve retornar falso para CPF com tamanho incorreto",
			input:    "123456",
			expected: false,
		},
		{
			name:     "Deve retornar falso para CPF com letras",
			input:    "abc.def.ghi-jk",
			expected: false,
		},
		{
			name:     "Deve retornar true para CPF Válido com Ponto e Traço ",
			input:    "529.982.247-25",
			expected: true,
		},
		{
			name:     "Deve retornar true para CPF Válido Somente Numeros",
			input:    "52998224725",
			expected: true,
		},
	}

	// Rodamos um loop passando por cada cenário da tabela
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := validator.IsCpf(tc.input)

			// Comparamos o resultado obtido com o esperado
			if result != tc.expected {
				t.Errorf("Para o input '%s', esperava %v, mas recebeu %v", tc.input, tc.expected, result)
			}
		})
	}
}
