package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/adrianjpsantos/rental-api/internal/security"
	"github.com/google/uuid"
)

// UserService contém toda a lógica de negócio relacionada a usuários
type UserService struct {
	repository user.Repository
}

func (s *UserService) GetUserForAuthentication(ctx context.Context, email string) (*user.UserForAuthentication, error) {
	return s.repository.GetUserForAuthentication(ctx, email)
}

func (s *UserService) Update(ctx context.Context, id uuid.UUID, input user.UserUpdateInput) error {
	fmt.Println("USER SERVICE UPDATE: Only call repository direct")
	return s.Update(ctx, id, input)
}

func NewUserService(repo user.Repository) user.Service {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) Create(ctx context.Context, input user.UserCreateInput) error {

	exists, err := s.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	if exists {
		return user.ErrEmailAlreadyExists
	}

	if input.Cpf != "" {
		exists, err = s.ExistsByCPF(ctx, input.Cpf)
		if err != nil {
			return err
		}
		if exists {
			return user.ErrCPFAlreadyExists
		}
	}

	passwordHashed, err := security.GenerateHashedPassword(input.Password)
	if err != nil {
		return err
	}

	input.Password = passwordHashed
	newUser, err := user.NewUser(input)

	if err != nil {
		return err
	}

	err = s.repository.Create(ctx, *newUser)
	if err != nil {
		return err
	}

	return nil
}

// GetByID busca um usuário pelo ID
func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	existingUser, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

// GetByEmail busca usuário pelo email (útil para login)
func (s *UserService) GetByEmail(ctx context.Context, email string) (*user.UserPublic, error) {
	existingUser, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

func (s *UserService) GetByCPF(ctx context.Context, cpf string) (*user.UserPublic, error) {
	existingUser, err := s.repository.GetByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

// UpdateProfile atualiza dados do perfil do usuário
func (s *UserService) UpdateProfile(ctx context.Context, id uuid.UUID, input user.UserUpdateInput) (*user.User, error) {

	// Busca o usuário atual
	existingUser, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}

	// Regra de negócio: só o próprio usuário pode editar seu perfil
	// (a verificação de autorização normalmente é feita no Handler)

	// Atualiza apenas os campos enviados
	existingUser.Update(input)

	// Persiste as alterações
	if err := s.repository.Update(ctx, *existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return s.repository.ExistsByEmail(ctx, email)
}

func (s *UserService) ExistsByCPF(ctx context.Context, cpf string) (bool, error) {
	return s.repository.ExistsByCPF(ctx, cpf)
}

// UpdateReputation atualiza a reputação do usuário (usado após uma review)
func (s *UserService) UpdateReputationCache(ctx context.Context, id uuid.UUID) error {
	return s.repository.UpdateReputationCache(ctx, id)
}

func (s *UserService) UpdateTotalRentalCache(ctx context.Context, id uuid.UUID) error {
	return s.repository.UpdateTotalRentalCache(ctx, id)
}

func (s *UserService) UpdateTotalItemsRentedCache(ctx context.Context, id uuid.UUID) error {
	return s.repository.UpdateTotalItemsRentedCache(ctx, id)
}

// ListUsers (exemplo - útil para admin)
func (s *UserService) ListUsers(ctx context.Context) ([]*user.User, error) {
	// Aqui você pode adicionar filtros depois
	return nil, errors.New("não implementado ainda")
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
