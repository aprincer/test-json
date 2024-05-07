package service

import (
	"context"
	"test-json/data/response"
)

type TableInfoService interface {
	FindAll(ctx context.Context) []response.TableInfoResponse
	FindByName(ctx context.Context, name string) []response.TableInfoResponse
}
