package main

import (
	"fmt"
	"test-json/helper"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main(){
	fmt.Printf("Hello Word!")
	routes:=httprouter.New()

	routes.GET("/",func(w http.ResponseWriter, r *http.Request,p httprouter.Params){
		fmt.Fprint(w,"Hello World")

	})
	server:=http.Server{Addr:"localhost:8080",Handler:routes}
	err:=server.ListenAndServe()
	helper.PanicIfError(err)

}