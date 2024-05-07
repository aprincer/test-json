package repository

import (
	"context"
	"test-json/model"
)

type TableInfoRepository interface {
	FindAll(ctx context.Context) []model.TableInfo
	FindByName(ctx context.Context, name string) []model.TableInfo
}
