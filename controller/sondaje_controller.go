package controller

import (
	"net/http"
	"strconv"
	"test-json/data/response"
	"test-json/helper"
	"test-json/service"

	"github.com/julienschmidt/httprouter"
)

type SondajeController struct {
	SondajeService service.SondajeService
}

func NewSondajeController(sondajeService service.SondajeService) *SondajeController {
	return &SondajeController{SondajeService: sondajeService}
}

func (sc *SondajeController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result := sc.SondajeService.FindAll(r.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (sc *SondajeController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	strId := params.ByName("sondajeId")
	sondajeId, err := strconv.Atoi(strId)
	helper.PanicIfError(err)
	result := sc.SondajeService.FindById(r.Context(), sondajeId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)

}
