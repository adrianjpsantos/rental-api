package user

import (
	"github.com/google/uuid"
	"time"
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

func NewUser(name, email, passwordHash string, cpf, phone string, birthDate time.Time, role Role) (*User, error) {
	user := &User{
		Id:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		CPF:          cpf,
		Phone:        phone,
		BirthDate:    birthDate,
		Role:         role,
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
