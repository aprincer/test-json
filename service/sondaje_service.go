package service

import (
	"context"
	"test-json/data/response"
)

type SondajeService interface {
	FindAll(ctx context.Context) []response.SondajeResponse
	FindById(ctx context.Context, id int) response.SondajeResponse
}
