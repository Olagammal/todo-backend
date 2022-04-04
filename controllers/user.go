package controllers

import (
	"encoding/json"
	"net/http"
	"todo/database"
	"todo/models"
)

type UserCreatedResponse struct {
	Id       uint   `json:"id"`
	Message	 string `json:"message"`
}

type ErrorMessage struct {
	Error	string `json:"error"`
}

type LoginCredentials struct {
	Email		string 	`json:"email"`
	Password	string	`json:"password"`
}

func Register(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var userData models.User
	json.NewDecoder(r.Body).Decode(&userData)
	result := database.DB.Create(&userData)
	if(result != nil){
		json.NewEncoder(w).Encode(&UserCreatedResponse{userData.Id,"user is created successfully"})
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var loginCredentials LoginCredentials
	var userData models.User
	json.NewDecoder(r.Body).Decode(&loginCredentials)
	result:=database.DB.Where("email=?",loginCredentials.Email).First(&userData)
	if(result.Error != nil){
		json.NewEncoder((w)).Encode(&ErrorMessage{result.Error.Error()})
		return
	}else{
		if(loginCredentials.Password != userData.Password){
			json.NewEncoder((w)).Encode(&ErrorMessage{"password does not match"})
			return
		}
		json.NewEncoder((w)).Encode(&UserCreatedResponse{userData.Id,"this is the user id"})
	}
}