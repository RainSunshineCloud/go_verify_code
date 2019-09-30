package main

import (
	"github.com/gorilla/mux"
	"verify_code/handler"
)


func GetRouter () *mux.Router {
	mux1 := mux.NewRouter()
	mux1.HandleFunc("/get", handler.Get).Methods("GET")
	mux1.HandleFunc("/verify", handler.Verify).Methods("POST")
	mux1.HandleFunc("/ping", handler.Ping).Methods("POST")


	return mux1;
}
