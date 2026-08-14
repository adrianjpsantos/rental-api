package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

// UserService contém toda a lógica de negócio relacionada a usuários
type UserService struct {
	Repository user.Repository
}

func (s *UserService) Update(ctx context.Context, id uuid.UUID, input user.User) error {
	fmt.Println("USER SERVICE UPDATE: Only call repository direct")
	return s.Update(ctx, id, input)
}

func NewUserService(repo user.Repository) user.Service {
	return &UserService{
		Repository: repo,
	}
}

func (s *UserService) Create(ctx context.Context, input user.Role) (*uuid.UUID, error) {

	id, err := s.Repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// GetByID busca um usuário pelo ID
func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	existingUser, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

// ListUsers (exemplo - útil para admin)
func (s *UserService) ListUsers(ctx context.Context) ([]*user.User, error) {
	// Aqui você pode adicionar filtros depois
	return nil, errors.New("não implementado ainda")
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.Repository.Delete(ctx, id)
}
