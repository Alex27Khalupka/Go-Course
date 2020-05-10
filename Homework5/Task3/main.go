package main

import (
	"Homework5/Task3/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetHandler).Methods(http.MethodGet)
	router.HandleFunc("/", handlers.PostHandler).Methods(http.MethodPost)
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
