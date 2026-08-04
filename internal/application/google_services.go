package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/google"
	"google.golang.org/api/idtoken"
)

type GoogleService struct {
	ClientID string
}

// ValidateIdToken implements [google.Service].
func (g *GoogleService) ValidateIdToken(ctx context.Context, googleAuthInput google.GoogleAuthInput) (*google.GoogleTokenPayload, error) {

	payload, err := idtoken.Validate(ctx, googleAuthInput.IDToken, g.ClientID)
	if err != nil {
		// Tratar erro de validação
	}

	// Extraindo dados com segurança em Go:
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	googleID, _ := payload.Claims["sub"].(string)
	picture, _ := payload.Claims["picture"].(string)
	emailVerified, _ := payload.Claims["email_verified"].(bool)

	// Importante: valide se o e-mail foi verificado pelo Google antes de criar o usuário
	if !emailVerified {
		return nil, google.ErrEmailNotVerified
	}

	return &google.GoogleTokenPayload{
		Email:         email,
		Name:          name,
		Sub:           googleID,
		Picture:       picture,
		EmailVerified: emailVerified,
	}, nil
}

func NewGoogleService(clientId string) google.Service {
	return &GoogleService{
		ClientID: clientId,
	}
}
