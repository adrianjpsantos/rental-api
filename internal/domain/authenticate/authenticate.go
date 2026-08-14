package authenticate

import (
	"time"

	"github.com/google/uuid"
)

type AuthenticateInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthenticatePayload struct {
	UserID uuid.UUID `json:"user_id"`
}

type RegisterInput struct {
	FirstName string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string    `json:"last_name" validate:"required,min=2,max=100"`
	CPF       string    `json:"cpf" validate:"required,cpf"`
	Phone     string    `json:"phone" validate:"required,min=10,max=11,numeric"`
	BirthDate time.Time `json:"birth_date" validate:"required,adult"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,pass_strength"`
}
