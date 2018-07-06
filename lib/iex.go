package pricefetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type IexFetcher struct {
	Names []string
}

type quote struct {
	LatestPrice float64 `json:"latestPrice"`
}

type stock struct {
	Quote quote `json:"quote"`
}

func (f IexFetcher) Run() ([]Security, error) {
	var ret []Security
	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf(
		"https://api.iextrading.com/1.0/stock/market/batch?types=quote&symbols=%s",
		strings.Join(f.Names, ",")))
	if err != nil {
		return ret, err
	}
	resp_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	var quotes map[string]stock
	err = json.Unmarshal(resp_bytes, &quotes)
	if err != nil {
		return ret, err
	}

	for k, v := range quotes {
		ret = append(ret, Security{strings.ToLower(k), v.Quote.LatestPrice})
	}
	return ret, err
}
