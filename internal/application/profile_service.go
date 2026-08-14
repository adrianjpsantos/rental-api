package application

import (
	"context"
	"github.com/adrianjpsantos/rental-api/internal/domain/profile"
	"github.com/google/uuid"
)

type ProfileService struct {
	Repository profile.Repository
}

// Create implements [profile.Service].
func (p *ProfileService) Create(ctx context.Context, input profile.CreateInput) error {
	panic("unimplemented")
}

// Delete implements [profile.Service].
func (p *ProfileService) Delete(ctx context.Context, userID uuid.UUID) error {
	panic("unimplemented")
}

// GetByUserID implements [profile.Service].
func (p *ProfileService) GetByUserID(ctx context.Context, userID uuid.UUID) (*profile.Profile, error) {
	panic("unimplemented")
}

// Update implements [profile.Service].
func (p *ProfileService) Update(ctx context.Context, userID uuid.UUID, input profile.UpdateInput) (*profile.Profile, error) {
	panic("unimplemented")
}

func NewProfileService(repo profile.Repository) profile.Service {
	return &ProfileService{
		Repository: repo,
	}
}
