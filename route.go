package main

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func route() {


	r:=mux.NewRouter()
	r.HandleFunc("/test",Createuserdata).Methods("POST")
	r.HandleFunc("/test1",sign).Methods("POST")
	
	r.HandleFunc("/test2",login).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", r))
	
	
}