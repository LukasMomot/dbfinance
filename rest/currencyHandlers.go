package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lukasmomot/dbfinance/services/currency"
)

type CurrentRateResponse struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type CurrencyConvertionResponse struct {
	From             string  `json:"from"`
	To               string  `json:"to"`
	Rate             float64 `json:"rate"`
	ConvertionResult float64 `json:"convertionResult"`
}

func CalculateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	from := vars["from"]
	to := vars["to"]
	amountStr := vars["amount"]
	amount, _ := strconv.ParseFloat(amountStr, 64)

	rate, value, err := currency.ConvertCurrency(from, to, amount)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Cannot process request")
		return
	}
	result := &CurrencyConvertionResponse{
		From:             from,
		To:               to,
		Rate:             rate,
		ConvertionResult: value,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rsp, _ := json.Marshal(result)
	w.Write(rsp)
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
