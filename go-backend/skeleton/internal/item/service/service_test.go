package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/domain"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/port"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/service"
)

var _ port.ItemRepository = (*stubRepo)(nil)

type stubRepo struct {
	items   []domain.Item
	saveErr error
}

func (s *stubRepo) CreateItem(_ context.Context, item domain.Item) error {
	if s.saveErr != nil {
		return s.saveErr
	}
	s.items = append(s.items, item)
	return nil
}

func (s *stubRepo) ListItems(_ context.Context) ([]domain.Item, error) {
	return s.items, nil
}

func TestCreateItem_Success(t *testing.T) {
	svc := service.NewItemService(&stubRepo{})
	item, err := svc.CreateItem(context.Background(), "widget")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(item.ID) == "" {
		t.Fatal("expected non-empty ID")
	}
	if string(item.Name) != "widget" {
		t.Fatalf("want name=widget, got %q", item.Name)
	}
}

func TestCreateItem_EmptyName(t *testing.T) {
	svc := service.NewItemService(&stubRepo{})
	_, err := svc.CreateItem(context.Background(), "")
	if !errors.Is(err, domain.ErrEmptyName) {
		t.Fatalf("want ErrEmptyName, got %v", err)
	}
}

func TestCreateItem_RepoError(t *testing.T) {
	repoErr := errors.New("storage unavailable")
	svc := service.NewItemService(&stubRepo{saveErr: repoErr})
	_, err := svc.CreateItem(context.Background(), "widget")
	if !errors.Is(err, repoErr) {
		t.Fatalf("want repo error, got %v", err)
	}
}

func TestListItems_ReturnsSaved(t *testing.T) {
	repo := &stubRepo{}
	svc := service.NewItemService(repo)

	if _, err := svc.CreateItem(context.Background(), "alpha"); err != nil {
		t.Fatalf("create: %v", err)
	}
	if _, err := svc.CreateItem(context.Background(), "beta"); err != nil {
		t.Fatalf("create: %v", err)
	}

	items, err := svc.ListItems(context.Background())
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d", len(items))
	}
}
