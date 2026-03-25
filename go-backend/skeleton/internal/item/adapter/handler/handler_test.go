package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/adapter/handler"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/adapter/memory"
	"github.com/${{ values.repoOwner }}/${{ values.name }}/internal/item/service"
)

func newRouter() http.Handler {
	repo := memory.NewRepository()
	svc := service.NewItemService(repo)
	h := handler.NewHandler(svc)
	r := chi.NewRouter()
	h.Register(r)
	return r
}

func TestListItems_Empty(t *testing.T) {
	w := httptest.NewRecorder()
	newRouter().ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/items", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestCreateItem(t *testing.T) {
	r := newRouter()

	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBufferString(`{"name":"widget"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("create: want 201, got %d", w.Code)
	}

	var body map[string]any
	json.NewDecoder(w.Body).Decode(&body)
	if body["id"] == "" {
		t.Fatal("expected non-empty id")
	}
	if body["name"] != "widget" {
		t.Fatalf("expected name=widget, got %v", body["name"])
	}
}

func TestCreateItem_EmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBufferString(`{"name":""}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}
