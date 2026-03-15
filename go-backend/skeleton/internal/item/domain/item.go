package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Domain errors — translated to HTTP status by the adapter layer.
var (
	ErrEmptyName    = errors.New("name cannot be empty")
	ErrItemNotFound = errors.New("item not found")
)

// ID is a value object for item identifiers.
// Typed to prevent passing a Name where an ID is expected.
type ID string

func NewID() ID {
	return ID(uuid.NewString())
}

// Name is a value object: immutable, self-validating, no zero value.
type Name string

func NewName(s string) (Name, error) {
	if s = strings.TrimSpace(s); s == "" {
		return "", ErrEmptyName
	}
	return Name(s), nil
}

// Item is the aggregate root.
type Item struct {
	ID        ID
	Name      Name
	CreatedAt time.Time
}

// NewItem is the domain constructor — enforces invariants at creation time.
func NewItem(name Name) Item {
	return Item{
		ID:        NewID(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}
