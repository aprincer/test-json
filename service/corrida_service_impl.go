package service

import (
	"context"
	"test-json/data/request"
	"test-json/data/response"
	"test-json/helper"
	"test-json/model"
	"test-json/repository"
)

type CorridaServiceImpl struct {
	CorridaRepository repository.CorridaRepository
}

func NewCorridaServiceImpl(corridaRepository repository.CorridaRepository) CorridaService {
	return &CorridaServiceImpl{CorridaRepository: corridaRepository}
}

func (c *CorridaServiceImpl) Create(ctx context.Context, request request.CorridaCreateRequest) {
	corrida := model.Corrida{
		Name: request.Name,
	}
	c.CorridaRepository.Save(ctx, corrida)
}

func (c *CorridaServiceImpl) Delete(ctx context.Context, corridaId int) {
	corrida, err := c.CorridaRepository.FindById(ctx, corridaId)
	helper.PanicIfError(err)
	c.CorridaRepository.Delete(ctx, corrida.Id)
}
func (c *CorridaServiceImpl) FindAll(ctx context.Context) []response.CorridaResponse {
	corridas := c.CorridaRepository.FindAll(ctx)

	var corridaResp []response.CorridaResponse
	for _, value := range corridas {
		corrida := response.CorridaResponse{Id: value.Id, Name: value.Name}
		corridaResp = append(corridaResp, corrida)
	}
	return corridaResp
}

func (c *CorridaServiceImpl) FindById(ctx context.Context, corridaId int) response.CorridaResponse {
	corrida, err := c.CorridaRepository.FindById(ctx, corridaId)
	helper.PanicIfError(err)
	return response.CorridaResponse(corrida)
}

func (c *CorridaServiceImpl) Update(ctx context.Context, request request.CorridaUpdateRequest) {
	corrida, err := c.CorridaRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	corrida.Name = request.Name
	c.CorridaRepository.Update(ctx, corrida)
}
