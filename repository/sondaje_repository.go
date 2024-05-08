package repository

import (
	"context"
	"test-json/model"
)

type SondajeRepository interface {
	FindAll(ctx context.Context) []model.Sondaje
	FindById(ctx context.Context, id int) (model.Sondaje, error)
}
