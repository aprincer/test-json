package repository

import (
	"context"
	"database/sql"
	"test-json/helper"
	"test-json/model"
)

type SondajeRepositoryImp struct {
	Db *sql.DB
}

func NewSondajeRepository(Db *sql.DB) SondajeRepository {
	return &SondajeRepositoryImp{Db: Db}
}

func (s *SondajeRepositoryImp) FindAll(ctx context.Context) []model.Sondaje {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := ""
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)

	defer result.Close()
	sondajes := make(map[int]*model.Sondaje)

	for result.Next() {
		var sondaje model.Sondaje
		var corrida model.Corrida
		errResult := result.Scan(
			&sondaje.Id,
			&sondaje.Id_proyecto,
			&sondaje.Nombre_sondaje,
			&sondaje.Fecha_inicio,
			&sondaje.Fecha_fin,
			&sondaje.Metros,
			&sondaje.Este,
			&sondaje.Norte,
			&sondaje.Azimut,
			&sondaje.Elevacion,
			&sondaje.Inclinacion,
			&sondaje.Id_geologo,
			&sondaje.Id_supervisor,
			&sondaje.S_uuid,
			&sondaje.Load_Date,
			&sondaje.Loaded_By,
			&sondaje.Modified_Date,
			&sondaje.Modified_By,
			&sondaje.Ts,
			&corrida.Id,
			&corrida.Id_sondaje,
			&corrida.Secuencia,
			&corrida.Id_litologia_general,
			&corrida.Cod_lit_general,
			&corrida.Id_litologia_especifica,
			&corrida.Cod_lit_especifica,
			&corrida.Longitud_triturada,
			&corrida.Id_tramo_triturado,
			&corrida.Cod_origen_tramo_triturado,
			&corrida.Id_resistencia,
			&corrida.Cod_resistencia,
			&corrida.Desde,
			&corrida.Hasta,
			&corrida.Id_cond_subterranea,
			&corrida.Cod_cond_subterranea,
			&corrida.Recuperado_m,
			&corrida.Rqd_m,
			&corrida.Num_fracturas,
			&corrida.Id_jn,
			&corrida.Cod_jn,
			&corrida.Id_ja,
			&corrida.Cod_ja,
			&corrida.Id_jr,
			&corrida.Cod_jr,
			&corrida.Estructura,
			&corrida.Surface_cond,
			&corrida.Gsi_aprox,
			&corrida.Comentario,
			&corrida.Id_abertura,
			&corrida.Id_rugosidad,
			&corrida.Id_relleno,
			&corrida.Id_meteorizacion,
			&corrida.C_uuid,
			&corrida.Suelo,
			&corrida.Load_Date,
			&corrida.Loaded_By,
			&corrida.Modified_Date,
			&corrida.Modified_By,
			&corrida.Ts,
		)

		if errResult != nil {
			helper.PanicIfError(errResult)
		}

		//Añadimos sondaje al Mapa
		if _, ok := sondajes[sondaje.Id]; !ok {
			sondajes[sondaje.Id] = &model.Sondaje{
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
			}
		}
		//Añadimos corrida al Sondaje
		sondajes[sondaje.Id].Corridas = append(sondajes[sondaje.Id].Corridas, model.Corrida{
			Id:                         corrida.Id,
			Id_sondaje:                 corrida.Id_sondaje,
			Secuencia:                  corrida.Secuencia,
			Id_litologia_general:       corrida.Id_litologia_general,
			Cod_lit_general:            corrida.Cod_lit_general,
			Id_litologia_especifica:    corrida.Id_litologia_especifica,
			Cod_lit_especifica:         corrida.Cod_lit_especifica,
			Longitud_triturada:         corrida.Longitud_triturada,
			Id_tramo_triturado:         corrida.Id_tramo_triturado,
			Cod_origen_tramo_triturado: corrida.Cod_origen_tramo_triturado,
			Id_resistencia:             corrida.Id_resistencia,
			Cod_resistencia:            corrida.Cod_resistencia,
			Desde:                      corrida.Desde,
			Hasta:                      corrida.Hasta,
			Id_cond_subterranea:        corrida.Id_cond_subterranea,
			Cod_cond_subterranea:       corrida.Cod_cond_subterranea,
			Recuperado_m:               corrida.Recuperado_m,
			Rqd_m:                      corrida.Rqd_m,
			Num_fracturas:              corrida.Num_fracturas,
			Id_jn:                      corrida.Id_jn,
			Cod_jn:                     corrida.Cod_jn,
			Id_ja:                      corrida.Id_ja,
			Cod_ja:                     corrida.Cod_ja,
			Id_jr:                      corrida.Id_jr,
			Cod_jr:                     corrida.Cod_jr,
			Estructura:                 corrida.Estructura,
			Surface_cond:               corrida.Surface_cond,
			Gsi_aprox:                  corrida.Gsi_aprox,
			Comentario:                 corrida.Comentario,
			Id_abertura:                corrida.Id_abertura,
			Id_rugosidad:               corrida.Id_rugosidad,
			Id_relleno:                 corrida.Id_relleno,
			Id_meteorizacion:           corrida.Id_meteorizacion,
			C_uuid:                     corrida.C_uuid,
			Suelo:                      corrida.Suelo,
			Load_Date:                  corrida.Load_Date,
			Loaded_By:                  corrida.Loaded_By,
			Modified_Date:              corrida.Modified_Date,
			Modified_By:                corrida.Modified_By,
			Ts:                         corrida.Ts,
		})
	}

	var data []model.Sondaje

	for _, s := range sondajes {
		sond := model.Sondaje{
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
		}
		data = append(data, sond)
	}

	return data
}
