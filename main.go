package main

import (
	"log"
	"net/http"
	"todo/routes"
)

func main(){
	allRoutes := routes.SetupRouter()
	log.Fatal(http.ListenAndServe("127.0.0.1:8000",allRoutes))
}