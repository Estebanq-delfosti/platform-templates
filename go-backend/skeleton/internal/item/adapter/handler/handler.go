package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/port"
)

// Handler is the HTTP driving adapter (primary/driving).
// Translates HTTP ↔ domain calls via the input port (port.ItemService).
// It never touches service or repository implementations directly.
type Handler struct {
	svc port.ItemService
}

func NewHandler(svc port.ItemService) *Handler {
	return &Handler{svc: svc}
}

// Register mounts the item routes on the given router.
func (h *Handler) Register(r chi.Router) {
	r.Get("/items", h.listItems)
	r.Post("/items", h.createItem)
}

func (h *Handler) listItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.ListItems(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errResponse("failed to list items"))
		return
	}
	writeJSON(w, http.StatusOK, newListResponse(items))
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var req createItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errResponse("invalid request body"))
		return
	}

	item, err := h.svc.CreateItem(r.Context(), req.Name)
	if err != nil {
		writeJSON(w, httpStatusFor(err), errResponse(err.Error()))
		return
	}

	writeJSON(w, http.StatusCreated, newItemResponse(item))
}
