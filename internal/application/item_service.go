package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/google/uuid"
)

type ItemService struct {
	itemRepo item.InterfaceItemRepository
}

func NewItemService(itemRepo item.InterfaceItemRepository) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
	}
}

func (s *ItemService) Register(ctx context.Context, createInput item.ItemCreateInput) (*item.Item, error) {

	// Cria a entidade usando o construtor do domínio
	newItem, err := item.NewItem(createInput)
	if err != nil {
		return nil, err
	}

	// Persiste no banco
	if err := s.itemRepo.Create(ctx, newItem); err != nil {
		return nil, err
	}

	return newItem, nil
}

func (s *ItemService) GetByID(ctx context.Context, id uuid.UUID) (*item.Item, error) {
	existingItem, err := s.itemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingItem == nil {
		return nil, item.ErrItemNotFound
	}
	return existingItem, nil
}

func (s *ItemService) ListByFilters(ctx context.Context, itemFilter item.ItemFilter) ([]*item.Item, error) {
	listItem, err := s.itemRepo.ListByFilters(ctx, itemFilter)
	if err != nil {
		return nil, err
	}
	if len(listItem) <= 0 {
		return nil, item.ErrItemNotFound
	}
	return listItem, nil
}

func (s *ItemService) Update(ctx context.Context, itemID uuid.UUID, updated *item.ItemUpdate) (*item.Item, error) {

	existingItem, err := s.itemRepo.GetByID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	if existingItem == nil {
		return nil, item.ErrItemNotFound
	}

	existingItem.Update(*updated)

	// Persiste as alterações
	if err := s.itemRepo.Update(ctx, existingItem); err != nil {
		return nil, err
	}

	return existingItem, nil
}

func (s *ItemService) Delete(ctx context.Context, itemID uuid.UUID) error {

	err := s.itemRepo.Delete(ctx, itemID)

	if err != nil {
		return err
	}
	return nil
}

func (s *ItemService) Exists(ctx context.Context, itemID uuid.UUID) (bool, error) {
	exists, err := s.itemRepo.Exists(ctx, itemID)

	if err != nil {
		return false, err
	}

	return exists, nil
}
