package lambdahandler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/config"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/lambdahandler"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/memory"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/service"
)

func newHandler() *lambdahandler.Handler {
	repo := memory.NewRepository()
	svc := service.NewItemService(repo)
	return lambdahandler.NewHandler(svc, &config.Config{LogLevel: "INFO", Environment: "test"})
}

func makeRequest(method, path, body string) events.APIGatewayV2HTTPRequest {
	return events.APIGatewayV2HTTPRequest{
		RawPath: path,
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: method,
			},
		},
		Body: body,
	}
}

func TestListItems_Empty(t *testing.T) {
	resp, err := newHandler().Handle(context.Background(), makeRequest(http.MethodGet, "/items", ""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("want 200, got %d", resp.StatusCode)
	}
}

func TestCreateItem(t *testing.T) {
	resp, err := newHandler().Handle(context.Background(), makeRequest(http.MethodPost, "/items", `{"name":"widget"}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("create: want 201, got %d", resp.StatusCode)
	}

	var body map[string]any
	json.Unmarshal([]byte(resp.Body), &body)
	if body["id"] == "" {
		t.Fatal("expected non-empty id")
	}
	if body["name"] != "widget" {
		t.Fatalf("expected name=widget, got %v", body["name"])
	}
}

func TestCreateItem_EmptyName(t *testing.T) {
	resp, err := newHandler().Handle(context.Background(), makeRequest(http.MethodPost, "/items", `{"name":""}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", resp.StatusCode)
	}
}

func TestUnknownRoute(t *testing.T) {
	resp, err := newHandler().Handle(context.Background(), makeRequest(http.MethodGet, "/unknown", ""))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("want 404, got %d", resp.StatusCode)
	}
}
