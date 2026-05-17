package user

import (
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

type UserUpdate struct {
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
func (u *User) Update(update UserUpdate) error {
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
