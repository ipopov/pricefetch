package pricefetch

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/antchfx/xmlquery"
)

type VanguardFund struct {
	Id   int
	Name string
}

func (fund VanguardFund) GetName() string {
	return fund.Name
}

func (fund VanguardFund) GetPrice() (float64, error) {
	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://personal.vanguard.com/us/FundsRSS?FundId=%d", fund.Id))
	if err != nil {
		return 0, err
	}

	rss, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return 0, err
	}

	price := xmlquery.FindOne(rss, "//item/title")
	if price == nil {
		return 0, errors.New(rss.OutputXML(true))
	}
	capture_price := regexp.MustCompile("^[^$]*\\$([^ ]+).*")
	matches := capture_price.FindSubmatch([]byte(price.InnerText()))
	if matches == nil {
		return 0, errors.New("Regexp failed to match.")
	}
	if len(matches) != 2 {
		return 0, errors.New("Regexp failed to match.")
	}
	return strconv.ParseFloat(string(matches[1]), 64)
}
