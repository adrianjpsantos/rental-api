package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/google/uuid"
)

type CategoryService struct {
	repository category.Repository
}

func NewCategoryService(repo category.Repository) category.Service {
	return &CategoryService{
		repository: repo,
	}
}

func (s *CategoryService) Create(ctx context.Context, input category.CategoryCreateInput) error {

	exists, err := s.repository.Exists(ctx, input.Name)
	if err != nil {
		return err
	}
	if exists {
		return category.ErrNameAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	NewCategory, err := category.NewCategory(input)
	if err != nil {
		return err
	}

	// Persiste no banco
	if err := s.repository.Create(ctx, *NewCategory); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetByID(ctx context.Context, id uuid.UUID) (*category.Category, error) {
	existingCategory, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, category.ErrCategoryNotFound
	}
	return existingCategory, nil
}

func (s *CategoryService) Update(ctx context.Context, id uuid.UUID, update category.CategoryUpdate) error {

	existingCategory, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return category.ErrCategoryNotFound
	}

	existingCategory.Update(update)

	// Persiste as alterações
	if err := s.repository.Update(ctx, *existingCategory); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) List(ctx context.Context) ([]*category.Category, error) {

	list, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *CategoryService) Delete(ctx context.Context, slotID uuid.UUID) error {
	err := s.repository.Delete(ctx, slotID)

	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) Exists(ctx context.Context, name string) (bool, error) {
	return s.repository.Exists(ctx, name)
}
