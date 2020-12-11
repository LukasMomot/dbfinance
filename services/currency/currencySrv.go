package currency

import "fmt"

// ConvertCurrency converts the currency
func ConvertCurrency(from string, to string, amount float64) float64 {
	fmt.Printf("From: %s, to: %s", from, to)
	return amount * 4
}
