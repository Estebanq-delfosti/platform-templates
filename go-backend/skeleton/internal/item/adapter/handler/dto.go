package handler

// Request / Response structs live in the adapter — domain entities have no JSON tags.

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/domain"
)

type createItemRequest struct {
	Name string `json:"name"`
}

type itemResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type listItemsResponse struct {
	Items []itemResponse `json:"items"`
	Total int            `json:"total"`
}

// httpStatusFor maps domain errors to HTTP status codes.
// Adding a new domain error: add one case here — no other file changes needed.
func httpStatusFor(err error) int {
	switch {
	case errors.Is(err, domain.ErrEmptyName):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func newItemResponse(i domain.Item) itemResponse {
	return itemResponse{ID: string(i.ID), Name: string(i.Name), CreatedAt: i.CreatedAt}
}

func newListResponse(items []domain.Item) listItemsResponse {
	resp := listItemsResponse{Items: make([]itemResponse, len(items)), Total: len(items)}
	for i, item := range items {
		resp.Items[i] = newItemResponse(item)
	}
	return resp
}

func errResponse(msg string) map[string]string { return map[string]string{"error": msg} }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
