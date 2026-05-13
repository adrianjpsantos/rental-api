package application

import (
	"context"
	"errors"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/user"
	"github.com/google/uuid"
)

// UserService contém toda a lógica de negócio relacionada a usuários
type UserService struct {
	userRepo user.InterfaceUserRepository
}

// NewUserService cria uma nova instância do serviço
func NewUserService(userRepo user.InterfaceUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register cria um novo usuário no sistema
func (s *UserService) Register(ctx context.Context, name, email, passwordHash, cpf, phone string, birthDate time.Time, role user.Role) (*user.User, error) {

	// Verifica se o email já existe
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, user.ErrEmailAlreadyExists
	}

	// Verifica se o CPF já existe (se informado)
	if cpf != "" {
		exists, err = s.userRepo.ExistsByCPF(ctx, cpf)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, user.ErrCPFAlreadyExists
		}
	}

	// Cria a entidade usando o construtor do domínio
	newUser, err := user.NewUser(name, email, passwordHash, cpf, phone, birthDate, role)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetByID busca um usuário pelo ID
func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

// GetByEmail busca usuário pelo email (útil para login)
func (s *UserService) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}
	return existingUser, nil
}

// UpdateProfile atualiza dados do perfil do usuário
func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, updated user.UserUpdate) (*user.User, error) {

	// Busca o usuário atual
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, user.ErrUserNotFound
	}

	// Regra de negócio: só o próprio usuário pode editar seu perfil
	// (a verificação de autorização normalmente é feita no Handler)

	// Atualiza apenas os campos enviados
	existingUser.Update(updated)

	// Persiste as alterações
	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return s.userRepo.ExistsByEmail(ctx, email)
}

func (s *UserService) ExistsByCPF(ctx context.Context, cpf string) (bool, error) {
	return s.userRepo.ExistsByCPF(ctx, cpf)
}

// UpdateReputation atualiza a reputação do usuário (usado após uma review)
func (s *UserService) UpdateReputation(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.UpdateReputationCache(ctx, userID)
}

func (s *UserService) UpdateTotalRentalCache(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.UpdateTotalRentalCache(ctx, userID)
}

func (s *UserService) UpdateTotalItemsRentedCache(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.UpdateTotalItemsRentedCache(ctx, userID)
}

// ListUsers (exemplo - útil para admin)
func (s *UserService) ListUsers(ctx context.Context) ([]*user.User, error) {
	// Aqui você pode adicionar filtros depois
	return nil, errors.New("não implementado ainda")
}
