package currency

import (
	"encoding/json"
	"fmt"
	conf "github.com/lukasmomot/dbfinance"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ConvertCurrency converts the currency
func ConvertCurrency(from string, to string, amount float64) float64 {
	fmt.Printf("From: %s, to: %s", from, to)
	apiKey := os.Getenv(conf.FixerAPIKey)
	fmt.Printf("FIXER key for api: %s", apiKey)
	return amount * 4
}

func GetCurrentRate() float64 {
	apiKey := os.Getenv(conf.FixerAPIKey)
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s&base=EUR&symbols=PLN", apiKey)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Declared an empty interface
	var result map[string]interface{}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		log.Fatal(err)
	}
	rates := result["rates"].(map[string]interface{})
	pln := rates["PLN"]

	fmt.Println(pln)

	return pln.(float64)
}
