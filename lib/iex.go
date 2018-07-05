package pricefetch

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/antchfx/jsonquery"
)

type IexStock struct {
	Name string
}

func (x IexStock) GetName() string {
	return x.Name
}

func (stock IexStock) GetPrice() (float64, error) {
	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/quote", stock.Name))
	if err != nil {
		return 0, err
	}

	quote, err := jsonquery.Parse(resp.Body)
	if err != nil {
		return 0, err
	}

	price := jsonquery.FindOne(quote, "//latestPrice")
	if price == nil {
		return 0, errors.New(quote.InnerText())
	}
	return strconv.ParseFloat(price.InnerText(), 32)
}
