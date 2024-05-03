package controller

import (
	"net/http"
	"strconv"
	service "test-json/Service"
	"test-json/data/request"
	"test-json/data/response"
	"test-json/helper"

	"github.com/julienschmidt/httprouter"
)

type CorridaController struct {
	CorridaService service.CorridaService
}

func NewCorridaController(corridaService service.CorridaService) *CorridaController {
	return &CorridaController{CorridaService: corridaService}
}

func (controller *CorridaController) Create(writer http.ResponseWriter, rq *http.Request, params httprouter.Params) {
	corridaCreateRequest := request.CorridaCreateRequest{}
	helper.ReadRequestBody(rq, &corridaCreateRequest)

	controller.CorridaService.Create(rq.Context(), corridaCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CorridaController) Update(writer http.ResponseWriter, rq *http.Request, params httprouter.Params) {
	corridaUpdateRequest := request.CorridaUpdateRequest{}
	helper.ReadRequestBody(rq, &corridaUpdateRequest)

	corridaId := params.ByName("corridaId")
	id, err := strconv.Atoi(corridaId)
	helper.PanicIfError(err)
	corridaUpdateRequest.Id = id
	controller.CorridaService.Update(rq.Context(), corridaUpdateRequest)

	controller.CorridaService.Update(rq.Context(), corridaUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *CorridaController) Delete(writer http.ResponseWriter, rq *http.Request, params httprouter.Params) {
	corridaId := params.ByName("corridaId")
	id, err := strconv.Atoi(corridaId)
	helper.PanicIfError(err)

	controller.CorridaService.Delete(rq.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CorridaController) FindAll(writer http.ResponseWriter, rq *http.Request, params httprouter.Params) {
	result := controller.CorridaService.FindAll(rq.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CorridaController) FindById(writer http.ResponseWriter, rq *http.Request, params httprouter.Params) {
	corridaId := params.ByName("corridaId")
	id, err := strconv.Atoi(corridaId)
	helper.PanicIfError(err)

	result := controller.CorridaService.FindById(rq.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}
