package currency

import (
	"fmt"
	conf "github.com/lukasmomot/dbfinance"
	"os"
)

// ConvertCurrency converts the currency
func ConvertCurrency(from string, to string, amount float64) float64 {
	fmt.Printf("From: %s, to: %s", from, to)
	apiKey := os.Getenv(conf.FixerAPIKey)
	fmt.Printf("FIXER key for api: %s", apiKey)
	return amount * 4
}
