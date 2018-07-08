package pricefetch

import (
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

type VanguardFund struct {
	Id   int
	Name string
}

type VanguardFetcher struct {
	Funds []VanguardFund
}

func (fund VanguardFund) get() (float64, error) {
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://personal.vanguard.com/us/FundsRSS?FundId=%d", fund.Id))
	if err != nil {
		return 0, err
	}
	var rss struct {
		Text []byte `xml:"channel>item>title"`
	}
	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = charset.NewReaderLabel
	if err = d.Decode(&rss); err != nil {
		return 0, err
	}
	capture_price := regexp.MustCompile(`^Price as of [0-9/]+: \$([^ ]+).*$`)
	matches := capture_price.FindSubmatch(rss.Text)
	if matches == nil || len(matches) != 2 {
		return 0, errors.New("Regexp failed to match.")
	}
	return strconv.ParseFloat(string(matches[1]), 64)
}

func (f VanguardFetcher) Run() ([]Security, error) {
	ret := make([]Security, len(f.Funds))
	errs := make([]error, len(f.Funds))
	var wg sync.WaitGroup
	wg.Add(len(f.Funds))
	for out_index, fund := range f.Funds {
		go func(out_index int, fund VanguardFund) {
			var price float64
			price, errs[out_index] = fund.get()
			ret[out_index] = Security{fund.Name, price}
			wg.Done()
		}(out_index, fund)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}
