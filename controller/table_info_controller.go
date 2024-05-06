package controller

import (
	"net/http"
	"test-json/data/response"
	"test-json/helper"
	"test-json/service"

	"github.com/julienschmidt/httprouter"
)

type TableInfoController struct {
	TableInfoService service.TableInfoService
}

func NewTableInfoController(tableInfoService service.TableInfoService) *TableInfoController {
	return &TableInfoController{TableInfoService: tableInfoService}
}

func (tc *TableInfoController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result := tc.TableInfoService.FindAll(r.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)
}
