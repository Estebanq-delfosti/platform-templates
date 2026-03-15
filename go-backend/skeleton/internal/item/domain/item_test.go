package domain_test

import (
	"testing"
	"time"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
)

func TestNewName_Valid(t *testing.T) {
	n, err := domain.NewName("widget")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(n) != "widget" {
		t.Fatalf("want 'widget', got %q", n)
	}
}

func TestNewName_Trim(t *testing.T) {
	n, err := domain.NewName("  widget  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(n) != "widget" {
		t.Fatalf("want trimmed 'widget', got %q", n)
	}
}

func TestNewName_Empty(t *testing.T) {
	_, err := domain.NewName("")
	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestNewName_Whitespace(t *testing.T) {
	_, err := domain.NewName("   ")
	if err == nil {
		t.Fatal("expected error for whitespace-only name")
	}
}

func TestNewID_NotEmpty(t *testing.T) {
	id := domain.NewID()
	if string(id) == "" {
		t.Fatal("expected non-empty ID")
	}
}

func TestNewID_Unique(t *testing.T) {
	id1 := domain.NewID()
	id2 := domain.NewID()
	if id1 == id2 {
		t.Fatal("expected unique IDs")
	}
}

func TestNewItem(t *testing.T) {
	name, _ := domain.NewName("widget")
	item := domain.NewItem(name)

	if string(item.ID) == "" {
		t.Fatal("expected non-empty ID")
	}
	if item.Name != name {
		t.Fatalf("want Name=%q, got %q", name, item.Name)
	}
	if item.CreatedAt.IsZero() {
		t.Fatal("expected non-zero CreatedAt")
	}
	if item.CreatedAt.After(time.Now().UTC()) {
		t.Fatal("CreatedAt is in the future")
	}
}
