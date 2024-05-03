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

	SQL := "SELECT id,name FROM corrida"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var corridas []model.Corrida

	for result.Next() {
		corrida := model.Corrida{}
		errResult := result.Scan(&corrida.Id, &corrida.Name)
		helper.PanicIfError(errResult)
		corridas = append(corridas, corrida)
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
		err := result.Scan(&corrida.Id, &corrida.Name)
		helper.PanicIfError(err)
		return corrida, nil
	} else {
		return corrida, errors.New("Corrida not found")
	}
}

func (c *CorridaRepositoryImpl) Save(ctx context.Context, corrida model.Corrida) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO corrida (name) VALUES($1)"
	_, err = tx.ExecContext(ctx, SQL, corrida.Name)
	helper.PanicIfError(err)

}

func (c *CorridaRepositoryImpl) Update(ctx context.Context, corrida model.Corrida) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE corrida set name=$1 WHERE id=$2"
	_, err = tx.ExecContext(ctx, SQL, corrida.Name, corrida.Id)
	helper.PanicIfError(err)

}
