package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukasmomot/dbfinance/services/currency"
)

func CalculateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	v := currency.ConvertCurrency("PLN", "USD", 2)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Calculate Currency %f", v)
}

func GetCurrencyRate(w http.ResponseWriter, r *http.Request) {
	rate := currency.GetCurrentRate()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get Currency rate %f", rate)
}
