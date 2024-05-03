package service

import (
	"context"
	"test-json/data/request"
	"test-json/data/response"
)

type CorridaService interface {
	Create(ctx context.Context, request request.CorridaCreateRequest)
	Update(ctx context.Context, request request.CorridaUpdateRequest)
	Delete(ctx context.Context, corridaID int)
	FindById(ctx context.Context, corridaID int) response.CorridaResponse
	FindAll(ctx context.Context) []response.CorridaResponse
}
