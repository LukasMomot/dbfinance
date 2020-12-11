package currency

import (
	"fmt"
	"os"
)

// ConvertCurrency converts the currency
func ConvertCurrency(from string, to string, amount float64) float64 {
	fmt.Printf("From: %s, to: %s", from, to)
	apiKey := os.Getenv("FIXER_API_KEY")
	fmt.Printf("FIXER key for api: %s", apiKey)
	return amount * 4
}
