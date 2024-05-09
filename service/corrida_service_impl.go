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
		Secuencia: request.Secuencia,
		//Secuencia: request.Secuencia,
	}
	c.CorridaRepository.Save(ctx, corrida)
}

func (c *CorridaServiceImpl) Delete(ctx context.Context, corridaId int) {
	//corrida, err := c.CorridaRepository.FindById(ctx, 0)
	//helper.PanicIfError(err)
	c.CorridaRepository.Delete(ctx, 0)
}
func (c *CorridaServiceImpl) FindAll(ctx context.Context) []response.CorridaResponse {
	corridas := c.CorridaRepository.FindAll(ctx)

	var corridaResp []response.CorridaResponse
	for _, value := range corridas {
		corrida := response.CorridaResponse{
			//Col_name: value.Col_name, Data_type: value.Data_type,
			Id:                         value.Id,
			Id_sondaje:                 value.Id_sondaje,
			Secuencia:                  value.Secuencia,
			Id_litologia_general:       value.Id_litologia_general,
			Cod_lit_general:            value.Cod_lit_general,
			Id_litologia_especifica:    value.Id_litologia_especifica,
			Cod_lit_especifica:         value.Cod_lit_especifica,
			Longitud_triturada:         value.Longitud_triturada,
			Id_tramo_triturado:         value.Id_tramo_triturado,
			Cod_origen_tramo_triturado: value.Cod_origen_tramo_triturado,
			Id_resistencia:             value.Id_resistencia,
			Cod_resistencia:            value.Cod_resistencia,
			Desde:                      value.Desde,
			Hasta:                      value.Hasta,
			Id_cond_subterranea:        value.Id_cond_subterranea,
			Cod_cond_subterranea:       value.Cod_cond_subterranea,
			Recuperado_m:               value.Recuperado_m,
			Rqd_m:                      value.Rqd_m,
			Num_fracturas:              value.Num_fracturas,
			Id_jn:                      value.Id_jn,
			Cod_jn:                     value.Cod_jn,
			Id_ja:                      value.Id_ja,
			Cod_ja:                     value.Cod_ja,
			Id_jr:                      value.Id_jr,
			Cod_jr:                     value.Cod_jr,
			Estructura:                 value.Estructura,
			Surface_cond:               value.Surface_cond,
			Gsi_aprox:                  value.Gsi_aprox,
			Comentario:                 value.Comentario,
			Id_abertura:                value.Id_abertura,
			Id_rugosidad:               value.Id_rugosidad,
			Id_relleno:                 value.Id_relleno,
			Id_meteorizacion:           value.Id_meteorizacion,
			C_uuid:                     value.C_uuid,
			Suelo:                      value.Suelo,
			Load_Date:                  value.Load_Date,
			Loaded_By:                  value.Loaded_By,
			Modified_Date:              value.Modified_Date,
			Modified_By:                value.Modified_By,
			Ts:                         value.Ts,
			Fracturas:                  value.Fracturas,
		}
		corridaResp = append(corridaResp, corrida)
	}
	return corridaResp
}

func (c *CorridaServiceImpl) FindById(ctx context.Context, corridaId int) response.CorridaResponse {
	corrida, err := c.CorridaRepository.FindById(ctx, corridaId)
	helper.PanicIfError(err)
	corridaResp := response.CorridaResponse{Id: corrida.Id, Secuencia: corrida.Secuencia}
	return response.CorridaResponse(corridaResp)
}

func (c *CorridaServiceImpl) Update(ctx context.Context, request request.CorridaUpdateRequest) {
	corrida, err := c.CorridaRepository.FindById(ctx, 0)
	helper.PanicIfError(err)

	corrida.Secuencia = request.Secuencia
	c.CorridaRepository.Update(ctx, corrida)
}
