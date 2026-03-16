package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyName    = errors.New("name cannot be empty")
	ErrItemNotFound = errors.New("item not found")
)

type ID string

func NewID() ID {
	return ID(uuid.NewString())
}

type Name string

func NewName(s string) (Name, error) {
	if s = strings.TrimSpace(s); s == "" {
		return "", ErrEmptyName
	}
	return Name(s), nil
}

type Item struct {
	ID        ID
	Name      Name
	CreatedAt time.Time
}

func NewItem(name Name) Item {
	return Item{
		ID:        NewID(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}
