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
	vars := mux.Vars(r)
	from := vars["from"]
	to := vars["to"]

	rate, err := currency.GetCurrentRate(from, to)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Cannot process request")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get Currency rate %f", rate)
}
