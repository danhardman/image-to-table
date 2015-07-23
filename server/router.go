package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router provides the http routing for the application
func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")
	r.HandleFunc("/Convert", Convert).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
