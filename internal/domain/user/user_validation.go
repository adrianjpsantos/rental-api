package user

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Validate valida as regras de negócio da entidade User
func (u *User) Validate() error {
	if len(u.Name) < 3 || len(u.Name) > 100 {
		return ErrInvalidName
	}
	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}
	if u.PasswordHash == "" {
		return ErrInvalidPasswordHash
	}
	if u.BirthDate.After(time.Now()) {
		return ErrInvalidBirthDate
	}
	if u.Role != Admin && u.Role != Lessor && u.Role != Lessee {
		return ErrInvalidRole
	}

	// Validação simples de CPF (pode ser melhorada com biblioteca)
	if u.CPF != "" && !IsValidCPF(u.CPF) {
		return ErrInvalidCPF
	}

	return nil
}

// Funções auxiliares de validação
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidCPF valida CPF brasileiro completo (com ou sem máscara)
func IsValidCPF(cpf string) bool {
	cpf = cleanCPF(cpf)

	// Deve ter exatamente 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais (CPFs inválidos conhecidos)
	if isAllDigitsEqual(cpf) {
		return false
	}

	// Calcula os dígitos verificadores
	d1 := calculateFirstVerifierDigit(cpf)
	d2 := calculateSecondVerifierDigit(cpf, d1)

	// Verifica se os dígitos calculados batem com os informados
	return cpf[9] == byte(d1+'0') && cpf[10] == byte(d2+'0')
}

// cleanCPF remove máscara e caracteres não numéricos
func cleanCPF(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, " ", "")
	return cpf
}

// isAllDigitsEqual verifica se todos os dígitos são iguais
func isAllDigitsEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// calculateFirstVerifierDigit calcula o primeiro dígito verificador
func calculateFirstVerifierDigit(cpf string) int {
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

// calculateSecondVerifierDigit calcula o segundo dígito verificador
func calculateSecondVerifierDigit(cpf string, firstDigit int) int {
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	sum += firstDigit * 2

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
