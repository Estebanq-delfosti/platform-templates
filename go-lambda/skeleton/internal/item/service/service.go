package service

import (
	"context"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/port"
)

var _ port.ItemService = (*ItemService)(nil)

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
