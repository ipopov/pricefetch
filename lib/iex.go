package pricefetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IexFetcher struct {
	Names []string
  ApiKey string
}

func (f IexFetcher) Run() ([]Security, error) {
	var ret []Security
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf(
"https://cloud.iexapis.com/stable/stock/market/batch?types=quote&symbols=%s&token=%s",
		strings.Join(f.Names, ","), f.ApiKey))
	if err != nil {
		return ret, err
	}
	var quotes map[string]struct {
		Quote struct {
			LatestPrice float64 `json:"latestPrice"`
		} `json:"quote"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return ret, err
	}

	for k, v := range quotes {
		ret = append(ret, Security{strings.ToLower(k), v.Quote.LatestPrice})
	}
	return ret, err
}
