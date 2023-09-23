package pricefetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PolygonFetcher struct {
	Names  []string
	ApiKey string
}

func (f PolygonFetcher) Run() ([]Security, error) {
	var ret []Security
	client := http.Client{}
	for _, security := range f.Names {
		resp, err := client.Get(fmt.Sprintf(
			"https://api.polygon.io/v2/aggs/ticker/%s/prev?adjusted=true&apiKey=%s",
			strings.ToUpper(security), f.ApiKey))
		if err != nil {
			return ret, err
		}
		var quote struct {
			Result []struct {
				ClosePrice float64 `json:"c"`
			} `json:"results"`
		}

		if err = json.NewDecoder(resp.Body).Decode(&quote); err != nil {
			return ret, err
		}

		ret = append(ret, Security{security, quote.Result[0].ClosePrice})
		// 5 calls per minute limit :-(
		time.Sleep(20 * time.Second)
	}
	return ret, nil
}
