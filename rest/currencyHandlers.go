package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukasmomot/dbfinance/services/currency"
)

type CurrentRateResponse struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

func CalculateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: IMPLEMENT THE METHOD
	vars := mux.Vars(r)
	fmt.Println(vars)

	v := currency.ConvertCurrency("PLN", "USD", 2)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Calculate Currency %f", v)
}

func GetCurrencyRate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	from := vars["from"]
	to := vars["to"]

	rate, err := currency.GetCurrentRate(from, to)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Cannot process request")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := &CurrentRateResponse{
		From: from,
		To:   to,
		Rate: rate,
	}

	rsp, _ := json.Marshal(response)
	w.Write(rsp)
}
