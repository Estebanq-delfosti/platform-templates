package port

// Repository ports (driven/secondary) — define WHAT the application needs from infrastructure.
// Implemented by: driven adapters (memory, postgres, redis, …).
// Consumed by: service layer.
//
//go:generate mockgen -source=repository.go -destination=mock/repository.go -package=mock

import (
	"context"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/domain"
)

// ItemRepository is the output port for item persistence.
type ItemRepository interface {
	CreateItem(ctx context.Context, item domain.Item) error
	ListItems(ctx context.Context) ([]domain.Item, error)
}
