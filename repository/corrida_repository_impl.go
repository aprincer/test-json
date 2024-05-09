package repository

import (
	"context"
	"database/sql"
	"errors"
	"test-json/helper"
	"test-json/model"
)

type CorridaRepositoryImpl struct {
	Db *sql.DB
}

func NewCorridaRepository(Db *sql.DB) CorridaRepository {
	return &CorridaRepositoryImpl{Db: Db}
}

func (c *CorridaRepositoryImpl) Delete(ctx context.Context, corridaID int) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	SQL := "DELETE FROM corrida WHERE id=$1"
	_, err1 := tx.ExecContext(ctx, SQL, corridaID)
	helper.PanicIfError(err1)
}

func (c *CorridaRepositoryImpl) FindAll(ctx context.Context) []model.Corrida {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//SQL := "SELECT * FROM corrida"
	//SQL := "describe table corrida"}
	//SQL := "select COLUMN_NAME, data_type FROM   INFORMATION_SCHEMA.Columns where table_name = 'corrida'"

	SQL := "SELECT * FROM corrida c LEFT JOIN fractura f ON c.id=f.id_corrida"

	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()
	corridaMap := make(map[int]*model.Corrida)

	//var corridas []model.Corrida

	for result.Next() {
		var corrida model.Corrida
		var fractura model.Fractura

		errResult := result.Scan(
			//&corrida.Col_name,
			//&corrida.Data_type,
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
			&fractura.Id,
			&fractura.Id_sondaje,
			&fractura.Id_corrida,
			&fractura.Id_tipo_discontinuidad,
			&fractura.Cod_tipo_discontinuidad,
			&fractura.Id_forma,
			&fractura.Cod_forma,
			&fractura.Id_tipo_relleno,
			&fractura.Cod_tipo_relleno,
			&fractura.Id_abertura,
			&fractura.Cod_abertura,
			&fractura.Id_rugosidad,
			&fractura.Cod_rugosidad,
			&fractura.Id_relleno,
			&fractura.Cod_relleno,
			&fractura.Id_meteorizacion,
			&fractura.Cod_meteorizacion,
			&fractura.Angulo_alfa,
			&fractura.Angulo_beta,
			&fractura.Profundidad,
			&fractura.Id_jr,
			&fractura.Id_ja,
			&fractura.Comentario,
			&fractura.Id_tipo_relleno2,
			&fractura.Angulo_gama,
			&fractura.F_uuid,
			&fractura.Load_Date,
			&fractura.Loaded_By,
			&fractura.Modified_Date,
			&fractura.Modified_By,
			&fractura.Ts,
		)
		helper.PanicIfError(errResult)
		//if errResult != nil {
		//corridas = append(corridas, corrida)
		//}
		if _, ok := corridaMap[corrida.Id]; !ok {
			corridaMap[corrida.Id] = &model.Corrida{
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
			}
		}
		corridaMap[corrida.Id].Fracturas = append(corridaMap[corrida.Id].Fracturas, model.Fractura{
			Id:                      fractura.Id,
			Id_sondaje:              fractura.Id,
			Id_corrida:              fractura.Id_corrida,
			Id_tipo_discontinuidad:  fractura.Id_tipo_discontinuidad,
			Cod_tipo_discontinuidad: fractura.Cod_tipo_discontinuidad,
			Id_forma:                fractura.Id_forma,
			Cod_forma:               fractura.Cod_forma,
			Id_tipo_relleno:         fractura.Id_tipo_relleno,
			Cod_tipo_relleno:        fractura.Cod_tipo_relleno,
			Id_abertura:             fractura.Id_abertura,
			Cod_abertura:            fractura.Cod_abertura,
			Id_rugosidad:            fractura.Id_rugosidad,
			Cod_rugosidad:           fractura.Cod_rugosidad,
			Id_relleno:              fractura.Id_relleno,
			Cod_relleno:             fractura.Cod_relleno,
			Id_meteorizacion:        fractura.Id_meteorizacion,
			Cod_meteorizacion:       fractura.Cod_meteorizacion,
			Angulo_alfa:             fractura.Angulo_alfa,
			Angulo_beta:             fractura.Angulo_beta,
			Profundidad:             fractura.Profundidad,
			Id_jr:                   fractura.Id_jr,
			Id_ja:                   fractura.Id_ja,
			Comentario:              fractura.Comentario,
			Id_tipo_relleno2:        fractura.Id_tipo_relleno2,
			Angulo_gama:             fractura.Angulo_gama,
			F_uuid:                  fractura.F_uuid,
			Load_Date:               fractura.Load_Date,
			Loaded_By:               fractura.Loaded_By,
			Modified_Date:           fractura.Modified_Date,
			Modified_By:             fractura.Modified_By,
			Ts:                      fractura.Ts,
		})
	}
	var data []model.Corrida

	for _, c := range corridaMap {
		cor := model.Corrida{
			Id:                         c.Id,
			Id_sondaje:                 c.Id_sondaje,
			Secuencia:                  c.Secuencia,
			Id_litologia_general:       c.Id_litologia_general,
			Cod_lit_general:            c.Cod_lit_general,
			Id_litologia_especifica:    c.Id_litologia_especifica,
			Cod_lit_especifica:         c.Cod_lit_especifica,
			Longitud_triturada:         c.Longitud_triturada,
			Id_tramo_triturado:         c.Id_tramo_triturado,
			Cod_origen_tramo_triturado: c.Cod_origen_tramo_triturado,
			Id_resistencia:             c.Id_resistencia,
			Cod_resistencia:            c.Cod_resistencia,
			Desde:                      c.Desde,
			Hasta:                      c.Hasta,
			Id_cond_subterranea:        c.Id_cond_subterranea,
			Cod_cond_subterranea:       c.Cod_cond_subterranea,
			Recuperado_m:               c.Recuperado_m,
			Rqd_m:                      c.Rqd_m,
			Num_fracturas:              c.Num_fracturas,
			Id_jn:                      c.Id_jn,
			Cod_jn:                     c.Cod_jn,
			Id_ja:                      c.Id_ja,
			Cod_ja:                     c.Cod_ja,
			Id_jr:                      c.Id_jr,
			Cod_jr:                     c.Cod_jr,
			Estructura:                 c.Estructura,
			Surface_cond:               c.Surface_cond,
			Gsi_aprox:                  c.Gsi_aprox,
			Comentario:                 c.Comentario,
			Id_abertura:                c.Id_abertura,
			Id_rugosidad:               c.Id_rugosidad,
			Id_relleno:                 c.Id_relleno,
			Id_meteorizacion:           c.Id_meteorizacion,
			C_uuid:                     c.C_uuid,
			Suelo:                      c.Suelo,
			Load_Date:                  c.Load_Date,
			Loaded_By:                  c.Loaded_By,
			Modified_Date:              c.Modified_Date,
			Modified_By:                c.Modified_By,
			Ts:                         c.Ts,
			Fracturas:                  c.Fracturas,
		}
		data = append(data, cor)
	}
	return data
}

func (c *CorridaRepositoryImpl) FindById(ctx context.Context, corridaID int) (model.Corrida, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id,secuencia FROM corrida WHERE id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, corridaID)
	helper.PanicIfError(errQuery)
	defer result.Close()

	corrida := model.Corrida{}

	if result.Next() {
		err := result.Scan(&corrida.Id, &corrida.Secuencia)
		helper.PanicIfError(err)
		return corrida, nil
	} else {
		return corrida, errors.New("corrida not found")
	}
}

func (c *CorridaRepositoryImpl) Save(ctx context.Context, corrida model.Corrida) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO corrida (name) VALUES($1)"
	_, err = tx.ExecContext(ctx, SQL, corrida.Id)
	helper.PanicIfError(err)

}

func (c *CorridaRepositoryImpl) Update(ctx context.Context, corrida model.Corrida) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE corrida set name=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, SQL, corrida.Id, corrida.Secuencia)
	helper.PanicIfError(err)

}
