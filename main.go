package main

import (
	"fmt"
	"net/http"
	service "test-json/Service"
	"test-json/config"
	"test-json/controller"
	"test-json/helper"
	"test-json/repository"
	"test-json/router"
)

func main() {
	fmt.Printf("Server running on port  %d!", 8080)
	//database
	db, errDb := config.DatabaseConnection()
	helper.PanicIfError(errDb)
	//repository
	corridaRepository := repository.NewCorridaRepository(db)

	//service
	corridaService := service.NewCorridaServiceImpl(corridaRepository)

	//controller
	corridaController := controller.NewCorridaController(corridaService)

	//router
	routes := router.NewRouter(corridaController)

	server := http.Server{Addr: "localhost:8080", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
