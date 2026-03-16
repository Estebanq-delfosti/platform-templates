package port

import (
	"context"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
)

type ItemCreator interface {
	CreateItem(ctx context.Context, name string) (domain.Item, error)
}

type ItemLister interface {
	ListItems(ctx context.Context) ([]domain.Item, error)
}

type ItemService interface {
	ItemCreator
	ItemLister
}
