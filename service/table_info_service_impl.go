package service

import (
	"context"
	"test-json/data/response"
	"test-json/repository"
)

type TableInfoServiceImpl struct {
	TableInfoRepository repository.TableInfoRepository
}

func NewTableInfoServiceImpl(tableInfoRepository repository.TableInfoRepository) TableInfoService {
	return &TableInfoServiceImpl{TableInfoRepository: tableInfoRepository}
}

func (ti *TableInfoServiceImpl) FindAll(ctx context.Context) []response.TableInfoResponse {
	tb := ti.TableInfoRepository.FindAll(ctx)
	var resp []response.TableInfoResponse
	for _, value := range tb {
		tbl := response.TableInfoResponse{
			Col_name: value.Col_name, Data_type: value.Data_type,
		}
		//fmt.Printf("data: %s", value.Col_name)
		resp = append(resp, tbl)
	}
	return resp
}

func (ti *TableInfoServiceImpl) FindByName(ctx context.Context, name string) []response.TableInfoResponse {
	list := ti.TableInfoRepository.FindByName(ctx, name)
	var resp []response.TableInfoResponse
	for _, value := range list {
		tbl := response.TableInfoResponse{
			Col_name:  value.Col_name,
			Data_type: value.Data_type,
		}
		resp = append(resp, tbl)
	}

	return resp
}
