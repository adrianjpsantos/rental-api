package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/google/uuid"
)

type CategoryService struct {
	categoryRepo category.InterfaceCategoryRepository
}

func NewCategoryService(categoryRepo category.InterfaceCategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) Register(ctx context.Context, name, description, icon string, position int, isActive bool) (*category.Category, error) {

	exists, err := s.categoryRepo.Exists(ctx, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, category.ErrNameAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	NewCategory, err := category.NewCategory(name, description, icon, position)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.categoryRepo.Create(ctx, NewCategory); err != nil {
		return nil, err
	}

	return NewCategory, nil
}

func (s *CategoryService) GetByID(ctx context.Context, categoryID uuid.UUID) (*category.Category, error) {
	existingCategory, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, category.ErrCategoryNotFound
	}
	return existingCategory, nil
}

func (s *CategoryService) Update(ctx context.Context, categoryID uuid.UUID, updated category.CategoryUpdate) (*category.Category, error) {

	existingCategory, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, category.ErrCategoryNotFound
	}

	existingCategory.Update(updated)

	// Persiste as alterações
	if err := s.categoryRepo.Update(ctx, existingCategory); err != nil {
		return nil, err
	}

	return existingCategory, nil
}

func (s *CategoryService) ListCategories(ctx context.Context) ([]*category.Category, error) {

	list, err := s.categoryRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *CategoryService) Delete(ctx context.Context, slotID uuid.UUID) error {
	err := s.categoryRepo.Delete(ctx, slotID)

	if err != nil {
		return err
	}

	return nil
}
