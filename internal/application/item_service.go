package application

import (
	"context"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/google/uuid"
)

type ItemService struct {
	repository item.Repository
}

func NewItemService(repo item.Repository) item.Service {
	return &ItemService{
		repository: repo,
	}
}

func (s *ItemService) Create(ctx context.Context, input item.ItemCreateInput) error {

	// Cria a entidade usando o construtor do domínio
	newItem, err := item.NewItem(input)
	if err != nil {
		return err
	}

	// Persiste no banco
	if err := s.repository.Create(ctx, newItem); err != nil {
		return err
	}

	return nil
}

func (s *ItemService) GetByID(ctx context.Context, id uuid.UUID) (*item.Item, error) {
	existingItem, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingItem == nil {
		return nil, item.ErrItemNotFound
	}
	return existingItem, nil
}

func (s *ItemService) ListByFilters(ctx context.Context, itemFilter item.ItemFilter) ([]*item.Item, error) {
	listItem, err := s.repository.ListByFilters(ctx, itemFilter)
	if err != nil {
		return nil, err
	}
	if len(listItem) <= 0 {
		return nil, item.ErrItemNotFound
	}
	return listItem, nil
}

func (s *ItemService) Update(ctx context.Context, itemID uuid.UUID, updated item.ItemUpdateInput) error {

	existingItem, err := s.repository.GetByID(ctx, itemID)
	if err != nil {
		return err
	}
	if existingItem == nil {
		return item.ErrItemNotFound
	}

	existingItem.Update(updated)

	// Persiste as alterações
	if err := s.repository.Update(ctx, existingItem); err != nil {
		return err
	}

	return nil
}

func (s *ItemService) Delete(ctx context.Context, itemID uuid.UUID) error {

	err := s.repository.Delete(ctx, itemID)

	if err != nil {
		return err
	}
	return nil
}

func (s *ItemService) Exists(ctx context.Context, itemID uuid.UUID) (bool, error) {
	exists, err := s.repository.Exists(ctx, itemID)

	if err != nil {
		return false, err
	}

	return exists, nil
}
