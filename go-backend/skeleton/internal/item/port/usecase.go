package port

// Use case ports (driving/primary) — define WHAT the application offers.
// Implemented by: service layer.
// Consumed by: driving adapters (HTTP, gRPC, CLI, …).
//
//go:generate mockgen -source=usecase.go -destination=mock/usecase.go -package=mock

import (
	"context"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
)

// ItemCreator is the use case for creating a new item.
type ItemCreator interface {
	CreateItem(ctx context.Context, name string) (domain.Item, error)
}

// ItemLister is the use case for listing all items.
type ItemLister interface {
	ListItems(ctx context.Context) ([]domain.Item, error)
}

// ItemService composes all item use cases.
// Depend on ItemCreator or ItemLister directly when a consumer only needs one operation;
// use ItemService when full access is required (e.g. the HTTP controller).
type ItemService interface {
	ItemCreator
	ItemLister
}
