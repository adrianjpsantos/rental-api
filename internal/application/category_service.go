package application

import (
	"context"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/category"
	"github.com/google/uuid"
)

type CategoryService struct {
	categorySlotRepo category.InterfaceCategoryRepository
}

func NewCategoryService(categorySlotRepo category.InterfaceCategoryRepository) *CategoryService {
	return &CategoryService{
		categorySlotRepo: categorySlotRepo,
	}
}

func (s *CategoryService) Register(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time, slotType category.CategoryType, reason category.Reason) (*category.CategorySlot, error) {

	exists, err := s.categorySlotRepo.Exists(name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, category.ErrSlotAlreadyExists
	}

	// Cria a entidade usando o construtor do domínio
	newSlot, err := category.NewCategorySlot(itemID, startDate, endDate, slotType, reason)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.categorySlotRepo.Create(newSlot); err != nil {
		return nil, err
	}

	return newSlot, nil
}

func (s *CategoryService) GetByID(ctx context.Context, slotID uuid.UUID) (*category.CategorySlot, error) {
	existingCategory, err := s.categorySlotRepo.GetByID(slotID)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, category.ErrSlotNotFound
	}
	return existingCategory, nil
}

func (s *CategoryService) FindByItemID(ctx context.Context, itemID uuid.UUID) ([]*category.CategorySlot, error) {
	listCategory, err := s.categorySlotRepo.FindByItemID(itemID)
	if err != nil {
		return nil, err
	}

	return listCategory, nil
}

func (s *CategoryService) FindOverlapping(ctx context.Context, itemID uuid.UUID, startDate, endDate time.Time) ([]*category.CategorySlot, error) {
	listCategory, err := s.categorySlotRepo.FindOverlapping(itemID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return listCategory, nil
}

func (s *CategoryService) Delete(ctx context.Context, slotID uuid.UUID) error {
	err := s.categorySlotRepo.Delete(slotID)

	if err != nil {
		return err
	}

	return nil
}
