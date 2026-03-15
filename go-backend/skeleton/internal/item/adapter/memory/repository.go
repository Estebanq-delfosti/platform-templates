package memory

import (
	"context"
	"sync"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/port"
)

// Compile-time check: Repository must satisfy port.ItemRepository (output port).
var _ port.ItemRepository = (*Repository)(nil)

// Repository is the in-memory driven adapter (secondary/driven).
// Swap for a postgres/redis adapter without touching domain, port, or service.
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
