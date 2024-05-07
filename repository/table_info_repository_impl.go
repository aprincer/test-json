package repository

import (
	"context"
	"database/sql"
	"test-json/helper"
	"test-json/model"
)

type TableInfoRepositoryImpl struct {
	Db *sql.DB
}

func NewTableInfoRepositoryImpl(db *sql.DB) TableInfoRepository {
	return &TableInfoRepositoryImpl{Db: db}
}

func (tr *TableInfoRepositoryImpl) FindAll(ctx context.Context) []model.TableInfo {
	tx, err := tr.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	SQL := "select COLUMN_NAME, data_type FROM   INFORMATION_SCHEMA.Columns where table_name = 'fractura'"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var tableInfo []model.TableInfo

	for result.Next() {
		tbl := model.TableInfo{}
		errResult := result.Scan(
			&tbl.Col_name,
			&tbl.Data_type,
		)
		//fmt.Printf("data: %s ", tbl.Col_name)
		if errResult == nil {
			tableInfo = append(tableInfo, tbl)
		} else {
			helper.PanicIfError(errResult)
		}
	}
	return tableInfo
}

func (tr *TableInfoRepositoryImpl) FindByName(ctx context.Context, name string) []model.TableInfo {
	tx, err := tr.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select CONCAT(COLUMN_NAME,' ',data_type) column_name, data_type FROM   INFORMATION_SCHEMA.Columns where table_name = $1"
	result, errQuery := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(errQuery)
	defer result.Close()
	var tableInfo []model.TableInfo

	for result.Next() {
		tbl := model.TableInfo{}
		errResult := result.Scan(
			&tbl.Col_name,
			&tbl.Data_type,
		)
		//fmt.Printf("data: %s ", tbl.Col_name)
		if errResult == nil {
			tableInfo = append(tableInfo, tbl)
		} else {
			helper.PanicIfError(errResult)
		}
	}
	return tableInfo
}
