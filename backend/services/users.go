package service

import (
	"context"
	"errors"
	"net/http"

	openapi "github.com/palledad/dManga/backend/openapi/gen"
)

// FindUsers -
func (s *ApiService) FindUsers(ctx context.Context, tags []string, limit int32) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), errors.New("FindUsers method not implemented")
}
