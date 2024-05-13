package router

import (
	"test-json/controller"

	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func NewRouter(corridaController *controller.CorridaController) *httprouter.Router {
func NewRouter(corridaController *controller.CorridaController, router *httprouter.Router) {
	//router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Welcome to my json-test API")
	})

	router.GET("/api/corrida", corridaController.FindAll)
	router.GET("/api/corrida/:corridaId", corridaController.FindById)
	router.GET("/api/corrida/:corridaId/:sondajeId", corridaController.FindBySondajeId)
	router.POST("/api/corrida", corridaController.Create)
	router.PATCH("/api/corrida/:corridaId", corridaController.Update)
	router.DELETE("/api/corrida/:corridaId", corridaController.Delete)
	//return router
}

// func NewRouterTableInfo(tbc *controller.TableInfoController) *httprouter.Router {
func NewRouterTableInfo(tbc *controller.TableInfoController, router *httprouter.Router) {
	//router := httprouter.New()
	router.GET("/api/table-info", tbc.FindAll)
	router.GET("/api/table-info/:name", tbc.FindByName)
}

func NewRouterSondaje(sc *controller.SondajeController, router *httprouter.Router) {
	router.GET("/api/sondaje", sc.FindAll)
	router.GET("/api/sondaje/:sondajeId", sc.FindById)
}
