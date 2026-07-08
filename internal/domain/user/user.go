package user

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin"
	Lessor Role = "lessor" // locador
	Lessee Role = "lessee" // locatário
)

type User struct {
	Id               uuid.UUID `json:"id,omitempty" validate:"required,uuid"`
	Name             string    `json:"name" validate:"min=8,max=100"`
	Email            string    `json:"email" validate:"required,email"`
	PasswordHash     string    `json:"password_hash,omitempty" validate:"required"`
	CPF              string    `json:"cpf" validate:"required,cpf"`
	Phone            string    `json:"phone" validate:"required,e164"`
	BirthDate        time.Time `json:"birth_date" validate:"required,datetime,adult"`
	AvatarURL        string    `json:"avatar_url,omitempty" validate:"url_encoded"`
	IsVerified       bool      `json:"is_verified" validate:"required,boolean"`
	Role             Role      `json:"role" validate:"required, role"`
	Reputation       float32   `json:"reputation" validate:"required,gte=0"`
	TotalRentals     int       `json:"total_rentals" validate:"required,gte=0"`
	TotalItemsRented int       `json:"total_items_rented" validate:"required,gte=0"`
	CreatedAt        time.Time `json:"created_at" validate:"datetime"`
	UpdatedAt        time.Time `json:"updated_at" validate:"datetime"`
	DeletedAt        time.Time `json:"deleted_at" validate:"datetime"` // Soft Delete
}

type UserUpdateInput struct {
	Name      string    `json:"name" validate:"min=8,max=100"`
	Phone     string    `json:"phone" validate:"required,e164"`
	AvatarURL string    `json:"avatar_url,omitempty" validate:"url_encoded"`
	BirthDate time.Time `json:"birth_date" validate:"required,datetime,adult"`
}

type UserForAuthentication struct {
	UserID       uuid.UUID `json:"user_id" validate:"required,uuid"`
	PasswordHash string    `json:"password_hash" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Name         string    `json:"name" validate:"min=8,max=100"`
}

type UserPublic struct {
	Id   uuid.UUID `json:"id,omitempty" validate:"required,uuid"`
	Name string    `json:"name" validate:"min=8,max=100"`
	Role Role      `json:"role" validate:"required, role"`
}

type UserCreateInput struct {
	Name      string    `json:"name" validate:"min=8,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,pass_strength"`
	Cpf       string    `json:"cpf" validate:"required,cpf"`
	Phone     string    `json:"phone" validate:"required,e164"`
	BirthDate time.Time `json:"birth_date" validate:"required,datetime,adult"`
	Role      Role      `json:"role" validate:"required, role"`
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
	validate := validator.New()

	if err := validate.Struct(u); err != nil {

		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
