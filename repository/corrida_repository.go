package repository

import (
	"context"
	"test-json/model"
)

type CorridaRepository interface {
	Save(ctx context.Context, corrida model.Corrida)
	Update(ctx context.Context, corrida model.Corrida)
	Delete(ctx context.Context, corridaId int)
	FindById(ctx context.Context, corridaId int) (model.Corrida, error)
	FindBySondajeId(ctx context.Context, sondajeId int) []model.Corrida
	FindAll(ctx context.Context) []model.Corrida
}
