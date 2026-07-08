package authenticate

import "github.com/google/uuid"

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
