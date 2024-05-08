package service

import (
	"context"
	"test-json/data/response"
	"test-json/helper"
	"test-json/repository"
)

type SondajeServiceImpl struct {
	SondajeRepository repository.SondajeRepository
}

func NewSondajeServiceImpl(sondajeRepository repository.SondajeRepository) SondajeService {
	return &SondajeServiceImpl{SondajeRepository: sondajeRepository}
}

func (s *SondajeServiceImpl) FindAll(ctx context.Context) []response.SondajeResponse {
	sondajes := s.SondajeRepository.FindAll(ctx)

	var data []response.SondajeResponse
	for _, sondaje := range sondajes {
		s := response.SondajeResponse{
			Id:             sondaje.Id,
			Id_proyecto:    sondaje.Id_proyecto,
			Nombre_sondaje: sondaje.Nombre_sondaje,
			Fecha_inicio:   sondaje.Fecha_inicio,
			Fecha_fin:      sondaje.Fecha_fin,
			Metros:         sondaje.Metros,
			Este:           sondaje.Este,
			Norte:          sondaje.Norte,
			Azimut:         sondaje.Azimut,
			Elevacion:      sondaje.Elevacion,
			Inclinacion:    sondaje.Inclinacion,
			Id_geologo:     sondaje.Id_geologo,
			Id_supervisor:  sondaje.Id_supervisor,
			S_uuid:         sondaje.S_uuid,
			Load_Date:      sondaje.Load_Date,
			Loaded_By:      sondaje.Loaded_By,
			Modified_Date:  sondaje.Modified_Date,
			Modified_By:    sondaje.Modified_By,
			Ts:             sondaje.Ts,
			Corridas:       sondaje.Corridas,
		}
		data = append(data, s)
	}
	return data
}

func (ss *SondajeServiceImpl) FindById(ctx context.Context, id int) response.SondajeResponse {
	s, err := ss.SondajeRepository.FindById(ctx, id)
	helper.PanicIfError(err)
	data := response.SondajeResponse{
		Id:             s.Id,
		Id_proyecto:    s.Id_proyecto,
		Nombre_sondaje: s.Nombre_sondaje,
		Fecha_inicio:   s.Fecha_inicio,
		Fecha_fin:      s.Fecha_fin,
		Metros:         s.Metros,
		Este:           s.Este,
		Norte:          s.Norte,
		Azimut:         s.Azimut,
		Elevacion:      s.Elevacion,
		Inclinacion:    s.Inclinacion,
		Id_geologo:     s.Id_geologo,
		Id_supervisor:  s.Id_supervisor,
		S_uuid:         s.S_uuid,
		Load_Date:      s.Load_Date,
		Loaded_By:      s.Loaded_By,
		Modified_Date:  s.Modified_Date,
		Modified_By:    s.Modified_By,
		Ts:             s.Ts,
		Corridas:       s.Corridas,
	}
	return response.SondajeResponse(data)
}
