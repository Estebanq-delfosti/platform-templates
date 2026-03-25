package service

import (
	"context"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/domain"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/port"
)

// Compile-time check: ItemService must satisfy port.ItemService (input port).
var _ port.ItemService = (*ItemService)(nil)

// ItemService is the application service.
// Implements the input port; depends only on the output port — never on concrete adapters.
type ItemService struct {
	repo port.ItemRepository
}

func NewItemService(repo port.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(ctx context.Context, name string) (domain.Item, error) {
	n, err := domain.NewName(name)
	if err != nil {
		return domain.Item{}, err
	}
	item := domain.NewItem(n)
	if err := s.repo.CreateItem(ctx, item); err != nil {
		return domain.Item{}, err
	}
	return item, nil
}

func (s *ItemService) ListItems(ctx context.Context) ([]domain.Item, error) {
	return s.repo.ListItems(ctx)
}
