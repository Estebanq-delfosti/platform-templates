package lambdahandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/config"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/port"
)

type Handler struct {
	svc    port.ItemService
	logger *slog.Logger
}

func NewHandler(svc port.ItemService, cfg *config.Config) *Handler {
	level := slog.LevelInfo
	if cfg.LogLevel == "DEBUG" {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	return &Handler{svc: svc, logger: logger}
}

func (h *Handler) Handle(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	method := req.RequestContext.HTTP.Method
	path := req.RawPath

	h.logger.InfoContext(ctx, "request", "method", method, "path", path)

	switch {
	case method == http.MethodGet && path == "/items":
		return h.listItems(ctx, req)
	case method == http.MethodPost && path == "/items":
		return h.createItem(ctx, req)
	default:
		return jsonResponse(http.StatusNotFound, errBody("route not found"))
	}
}

func (h *Handler) listItems(ctx context.Context, _ events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	items, err := h.svc.ListItems(ctx)
	if err != nil {
		h.logger.ErrorContext(ctx, "list items failed", "error", err)
		return jsonResponse(http.StatusInternalServerError, errBody("failed to list items"))
	}
	return jsonResponse(http.StatusOK, newListResponse(items))
}

func (h *Handler) createItem(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var body createItemRequest
	if err := json.Unmarshal([]byte(req.Body), &body); err != nil {
		return jsonResponse(http.StatusBadRequest, errBody("invalid request body"))
	}

	item, err := h.svc.CreateItem(ctx, body.Name)
	if err != nil {
		return jsonResponse(httpStatusFor(err), errBody(err.Error()))
	}
	return jsonResponse(http.StatusCreated, newItemResponse(item))
}
