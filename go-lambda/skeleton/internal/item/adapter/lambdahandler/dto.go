package lambdahandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"

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

func errBody(msg string) map[string]string { return map[string]string{"error": msg} }

func jsonResponse(status int, v any) (events.APIGatewayV2HTTPResponse, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusInternalServerError}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(b),
	}, nil
}
