package google

import "context"

type Service interface {
	ValidateIdToken(ctx context.Context, googleAuthInput GoogleAuthInput) (*GoogleTokenPayload, error)
}
