package port

import (
	"context"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/domain"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, item domain.Item) error
	ListItems(ctx context.Context) ([]domain.Item, error)
}
