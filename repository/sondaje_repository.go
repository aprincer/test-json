package repository

import (
	"context"
	"test-json/model"
)

type SondajeRepository interface {
	FindAll(ctx context.Context) []model.Sondaje
}
