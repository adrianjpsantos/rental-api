package authenticate

type AuthenticateInput struct {
	Email    string
	Password string
}

type AuthenticateOutput struct {
	Token        string
	RefreshToken string
}

type AuthenticatePayload struct {
	UserID string
	Email  string
}
