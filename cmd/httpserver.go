package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lukasmomot/dbfinance/rest"

	"github.com/gorilla/mux"
)

func main() {
	var err error
	fmt.Println("Digital Broker finance microservice is starting on port 3002")
	// Load configuration
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HealthHandler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	// Currency
	r.HandleFunc("/currency", rest.GetCurrencyRate)
	r.HandleFunc("/currency/convert", rest.CalculateCurrencyHandler).
		Queries(
			"from", "{from}",
			"to", "{to}",
			"amount", "{amount}").
		Methods("GET")

	err = http.ListenAndServe(":3002", r)
	if err != nil {
		log.Fatal("Cannot start http server...")
	}
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Digital Broker Finance Microservice is up and running")
}
