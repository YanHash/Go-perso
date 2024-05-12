package main

import (
	"fmt"
	"log"
	"net/http"

	"Go-perso/handlers"

	"github.com/gorilla/mux"
)

var creation bool = false

func main() {
	if !creation {
		var collections []string = []string{"Tools", "Education", "Experience", "Contacts"}
		handlers.CreateCollections(collections)
		creation = true
	}

	r := mux.NewRouter()

	handlers.InitMongoDB()
	r.HandleFunc("/tools", handlers.GetTools).Methods("GET")
	r.HandleFunc("/education", handlers.GetEducation).Methods("GET")
	r.HandleFunc("/experience", handlers.GetExperience).Methods("GET")
	r.HandleFunc("/contact", handlers.GetContact).Methods("GET")

	fmt.Printf("starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
