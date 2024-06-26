package main

import (
	"fmt"
	"net/http"
	"test-json/config"
	"test-json/controller"
	"test-json/helper"
	"test-json/repository"
	"test-json/router"
	"test-json/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Server running on port  %d!", 8080)
	//database
	db, errDb := config.DatabaseConnection()
	helper.PanicIfError(errDb)

	//repository
	corridaRepository := repository.NewCorridaRepository(db)
	tblRepository := repository.NewTableInfoRepositoryImpl(db)
	sondRepository := repository.NewSondajeRepository(db)
	//service
	corridaService := service.NewCorridaServiceImpl(corridaRepository)
	tblService := service.NewTableInfoServiceImpl(tblRepository)
	sondService := service.NewSondajeServiceImpl(sondRepository)
	//controller
	corridaController := controller.NewCorridaController(corridaService)
	tbController := controller.NewTableInfoController(tblService)
	sondController := controller.NewSondajeController(sondService)
	//router
	//routes := router.NewRouter(corridaController)
	routes := httprouter.New()
	router.NewRouter(corridaController, routes)
	router.NewRouterTableInfo(tbController, routes)
	router.NewRouterSondaje(sondController, routes)
	server := http.Server{Addr: "localhost:8080", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
