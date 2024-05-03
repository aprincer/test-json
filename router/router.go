package router

import (
	"test-json/controller"

	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(corridaController *controller.CorridaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Welcome to my json-test API")
	})

	router.GET("/api/corrida", corridaController.FindAll)
	router.GET("/api/corrida/:corridaId", corridaController.FindById)
	router.POST("/api/corrida", corridaController.Create)
	router.PATCH("/api/corrida/:corridaId", corridaController.Update)
	router.DELETE("/api/corrida/:corridaId", corridaController.Delete)
	return router
}
