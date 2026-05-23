package user

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin"
	Lessor Role = "lessor" // locador
	Lessee Role = "lessee" // locatário
)

type User struct {
	Id               uuid.UUID
	Name             string
	Email            string
	PasswordHash     string
	CPF              string
	Phone            string
	BirthDate        time.Time
	AvatarURL        string
	IsVerified       bool
	Role             Role
	Reputation       float32
	TotalRentals     int
	TotalItemsRented int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time // Soft Delete
}

type UserUpdateInput struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	AvatarURL string    `json:"avatar_url"`
	BirthDate time.Time `json:"birth_date"`
}

type UserForAuthentication struct {
	UserID       uuid.UUID `json:"user_id"`
	PasswordHash string    `json:"password_hash"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
}

type UserPublic struct {
	Id   uuid.UUID
	Name string
	Role Role
}

type UserCreateInput struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	Cpf       string    `json:"cpf" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
	Role      Role      `json:"role" validate:"required,oneof=admin lessor lessee"`
}

func NewUser(newUser UserCreateInput) (*User, error) {
	user := &User{
		Id:           uuid.New(),
		Name:         newUser.Name,
		Email:        newUser.Email,
		PasswordHash: newUser.Password, // A senha será hashada no serviço ("Recebe a senha em texto plano, mas armazena o hash ao criar a entidade")
		CPF:          newUser.Cpf,
		Phone:        newUser.Phone,
		BirthDate:    newUser.BirthDate,
		Role:         newUser.Role,
		IsVerified:   false,
		Reputation:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Update atualiza informações do usuário
func (u *User) Update(update UserUpdateInput) error {
	updated := false

	if update.Name != "" {
		u.Name = update.Name
		updated = true
	}
	if update.Phone != "" {
		u.Phone = update.Phone
		updated = true
	}
	if update.AvatarURL != "" {
		u.AvatarURL = update.AvatarURL
		updated = true
	}
	if !update.BirthDate.IsZero() {
		u.BirthDate = update.BirthDate
		updated = true
	}

	if updated {
		u.UpdatedAt = time.Now()
	}

	return u.Validate()
}

// AddReputation adiciona pontos à reputação (média)
func (u *User) AddReputation(newRating float32) {
	// Lógica simples de média (pode ser melhorada depois)
	if u.Reputation == 0 {
		u.Reputation = newRating
	} else {
		u.Reputation = (u.Reputation + newRating) / 2
	}
}

// IsLessor retorna se o usuário é locador
func (u *User) IsLessor() bool {
	return u.Role == Lessor || u.Role == Admin
}

// IsLessee retorna se o usuário é locatário
func (u *User) IsLessee() bool {
	return u.Role == Lessee || u.Role == Admin
}

// IsAdmin retorna se o usuário é administrador
func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

// CanCreateItem verifica se o usuário pode cadastrar itens para alugar
func (u *User) CanCreateItem() bool {
	return u.IsLessor() || u.IsAdmin()
}

// CanRentItem verifica se o usuário pode alugar itens
func (u *User) CanRentItem() bool {
	return u.IsLessee() || u.IsAdmin()
}

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
