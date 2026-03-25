package memory

import (
	"context"
	"sync"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/domain"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/port"
)

var _ port.ItemRepository = (*Repository)(nil)

type Repository struct {
	mu    sync.RWMutex
	store []domain.Item
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateItem(_ context.Context, item domain.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store = append(r.store, item)
	return nil
}

func (r *Repository) ListItems(_ context.Context) ([]domain.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]domain.Item{}, r.store...), nil
}
