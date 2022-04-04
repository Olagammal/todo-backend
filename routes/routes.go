package routes

import (
	"encoding/json"
	"net/http"

	"todo/controllers"

	"github.com/gorilla/mux"
)

func mainRoute(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode("hello world")
}

func SetupRouter() *mux.Router{
	r:=mux.NewRouter()
	r.HandleFunc("/",mainRoute).Methods("GET")
	r.HandleFunc("/register",controllers.Register).Methods("POST")
	r.HandleFunc("/login",controllers.Login).Methods("POST")

	return r
}