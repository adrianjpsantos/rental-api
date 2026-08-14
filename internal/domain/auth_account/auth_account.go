package authaccount

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Provider string

const (
	ProviderLocal     Provider = "local"
	ProviderGoogle    Provider = "google"
	ProviderFacebook  Provider = "facebook"
	ProviderGithub    Provider = "github"
	ProviderApple     Provider = "apple"
	ProviderMicrosoft Provider = "microsoft"
)

type AuthAccount struct {
	ID             uuid.UUID  `json:"id" validate:"required,uuid"`
	UserID         uuid.UUID  `json:"user_id" validate:"required,uuid"`
	Provider       Provider   `json:"provider" validate:"required,provider"`
	ProviderUserID string     `json:"provider_user_id" validate:"required"`
	Email          string     `json:"email" validate:"omitempty,email"`
	PasswordHash   *string    `json:"password_hash,omitempty"`
	EmailVerified  bool       `json:"email_verified"`
	IsPrimary      bool       `json:"is_primary"`
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`
	LinkedAt       time.Time  `json:"linked_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (a *AuthAccount) Validate() error {
	validate := validator.New()

	if err := validate.Struct(a); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

type RegisterLocalInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateInput struct {
	UserID         uuid.UUID `json:"user_id" validate:"required,uuid"`
	Provider       Provider  `json:"provider" validate:"required"`
	ProviderUserID string    `json:"provider_user_id" validate:"required"`
	Email          string    `json:"email" validate:"omitempty,email"`
	PasswordHash   *string   `json:"password_hash,omitempty"`
	IsPrimary      bool      `json:"is_primary"`
}

type UpdateInput struct {
	Email         string  `json:"email,omitempty" validate:"omitempty,email"`
	PasswordHash  *string `json:"password_hash,omitempty"`
	EmailVerified *bool   `json:"email_verified,omitempty"`
	IsPrimary     *bool   `json:"is_primary,omitempty"`
}

type Public struct {
	ID            uuid.UUID  `json:"id"`
	UserID        uuid.UUID  `json:"user_id"`
	Provider      Provider   `json:"provider"`
	Email         string     `json:"email,omitempty"`
	EmailVerified bool       `json:"email_verified"`
	IsPrimary     bool       `json:"is_primary"`
	LastLoginAt   *time.Time `json:"last_login_at,omitempty"`
}
