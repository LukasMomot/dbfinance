package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	conf "github.com/lukasmomot/dbfinance"
)

// ConvertCurrency converts the currency
func ConvertCurrency(from string, to string, amount float64) (float64, float64, error) {
	rate, err := GetCurrentRate(from, to)
	return rate, amount * rate, err
}

// GetCurrentRate gets the current currency rate
func GetCurrentRate(from string, to string) (float64, error) {
	apiKey := os.Getenv(conf.FixerAPIKey)
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s&base=%s&symbols=%s", apiKey, from, to)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Cannot perform request to fixer.io")
		fmt.Print(err.Error())
	}

	// Declared an empty interface
	var result map[string]interface{}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Cannot deserialize response")
	}

	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		log.Fatal("Cannot parse JSON response.")
	}

	if result["success"] == false {
		fmt.Println(result)
		return 0, fmt.Errorf("Request error")
	}

	rates := result["rates"].(map[string]interface{})
	rate := rates[to]

	fmt.Println(rate)

	return rate.(float64), nil
}
