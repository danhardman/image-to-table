package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Router provides the http routing for the application
func Router() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")
	r.HandleFunc("/Convert", Convert).Methods("POST")

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("public/static")))
	r.PathPrefix("/static/").Handler(staticHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Println("Listening to port: " + port)
}
