package profile

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	UserID    uuid.UUID
	FirstName string
	LastName  string
	CPF       string
	Phone     string
	BirthDate time.Time
	AvatarURL string
}

type CreateInput struct {
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid"`
	FirstName string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"omitempty,min=2,max=50"`
	CPF       string    `json:"cpf" validate:"required,cpf"`
	Phone     string    `json:"phone" validate:"required,e164"`
	BirthDate time.Time `json:"birth_date" validate:"required,adult"`
	AvatarURL string    `json:"avatar_url,omitempty" validate:"omitempty,url"`
}

type UpdateInput struct {
	FirstName string    `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"omitempty,min=2,max=50"`
	Phone     string    `json:"phone" validate:"omitempty,e164"`
	BirthDate time.Time `json:"birth_date" validate:"omitempty,adult"`
	AvatarURL string    `json:"avatar_url,omitempty" validate:"omitempty,url"`
}

type Public struct {
	UserID     uuid.UUID `json:"user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	AvatarURL  string    `json:"avatar_url,omitempty"`
	Reputation float32   `json:"reputation"`
}

func New(userID uuid.UUID, input CreateInput) (*Profile, error) {
	profile := &Profile{
		UserID:    userID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CPF:       input.CPF,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
		AvatarURL: input.AvatarURL,
	}

	return profile, nil
}
