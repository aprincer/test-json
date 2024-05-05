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
	SQL := "SELECT * FROM corrida"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var corridas []model.Corrida

	for result.Next() {
		corrida := model.Corrida{}

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
		)
		//helper.PanicIfError(errResult)
		if errResult != nil {
			corridas = append(corridas, corrida)
		}

	}
	return corridas
}

func (c *CorridaRepositoryImpl) FindById(ctx context.Context, corridaID int) (model.Corrida, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id,name FROM corrida WHERE id=$1"
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
