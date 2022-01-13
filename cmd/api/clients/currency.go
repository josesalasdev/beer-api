package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/josesalasdev/beer-api/cmd/api/config"
)

type Currency interface {
	Converter(from string, to string) (float64, error)
	Calculate(price, quantity float64, from, to string) (float64, error)
}

type currencyBody struct {
	Success bool               `json:"success"`
	Quotes  map[string]float64 `json:"quotes"`
}

type currency struct{}

func NewCurrencyClient() Currency {
	return &currency{}
}

// Converter is a method that get conversion between currencies.
func (c *currency) Converter(from string, to string) (float64, error) {
	path := fmt.Sprintf(
		"%v?access_key=%v&source=%v&currencies=%v",
		config.CurrencyPath,
		config.CurrencyAPISecret,
		strings.ToUpper(from),
		strings.ToUpper(to),
	)
	resp, err := http.Get(path)
	if err != nil {
		return 0, err
	}
	//We Read the response body on the line below.
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var body currencyBody
	if err := json.Unmarshal(bodyByte, &body); err != nil {
		return 0, err
	}
	if !body.Success {
		return 0, fmt.Errorf("Currencies not found !")
	}
	currencyJoin := strings.ToUpper(fmt.Sprintf("%v%v", from, to))
	return body.Quotes[currencyJoin], nil
}

// Converter is a method that get calculate the conversion between currencies.
func (c *currency) Calculate(price, quantity float64, from, to string) (float64, error) {
	if from == to {
		return (price * quantity), nil
	}
	priceConv, err := c.Converter(from, to)
	if err != nil {
		return 0, err
	}
	priceUnit := price / priceConv
	return (priceUnit * quantity), nil
}
