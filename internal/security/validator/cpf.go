package validator

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidateCpf(fl validator.FieldLevel) bool {
	return IsCpf(fl.Field().String())
}

func IsCpf(cpf string) bool {
	stripRegex := regexp.MustCompile(`\D`)

	cleanCpf := stripRegex.ReplaceAllString(cpf, "")

	// 2. Validações básicas (tamanho e sequências inválidas)
	if len(cleanCpf) != 11 || IsInvalidSequence(cleanCpf) {
		return false
	}

	digit1 := CalculateDigit(cleanCpf[:9], 10)
	digit2 := CalculateDigit(cleanCpf[:9+1], 11)

	d1, _ := strconv.Atoi(string(cleanCpf[9]))
	d2, _ := strconv.Atoi(string(cleanCpf[10]))

	return digit1 == d1 && digit2 == d2
}

func CalculateDigit(slice string, factor int) int {
	sum := 0
	for _, char := range slice {
		digit, _ := strconv.Atoi(string(char))
		sum += digit * factor
		factor--
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

func IsInvalidSequence(cpf string) bool {
	// Verifica sequências como 111.111.111-11
	// Em Go, uma forma eficiente é comparar se todos os bytes são iguais
	first := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != first {
			return false
		}
	}
	return true
}
