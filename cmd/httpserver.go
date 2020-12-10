package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Digital Broker finance microservice is starting on port 3002")
	r := mux.NewRouter()
	r.HandleFunc("/", HealthHandler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	err := http.ListenAndServe(":3002", r)
	if err != nil {
		log.Fatal("Cannot start http server...")
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Digital Broker Finance Microservice is up and running")
}
