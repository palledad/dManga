package service

import (
	openapi "github.com/palledad/dManga/backend/openapi/gen"
)

type ApiService struct {
}

// NewApiService creates an api service
func NewApiService() openapi.DefaultApiServicer {
	return &ApiService{}
}
