package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/database"
	"todo/routes"
)

func main(){
	allRoutes := routes.SetupRouter()
	connection:= database.ConnectToDatabase()	
	fmt.Println(connection)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000",allRoutes))
}