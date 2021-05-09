package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sync"
	"time"

	pf "github.com/ipopov/pricefetch/lib"
)

type Config struct {
	V pf.VanguardFetcher
	X pf.IexFetcher
}

func longestName(xs []pf.Security) int {
	ret := 0
	for _, x := range xs {
		if len(x.Name) > ret {
			ret = len(x.Name)
		}
	}
	return ret
}

func longestPrice(xs []pf.Security) int {
	ret := 0
	for _, x := range xs {
		// 4 is for the dollar sign, dot, cents.
		digits := 4 + 1 + int(math.Log10(x.Price))
		if digits > ret {
			ret = digits
		}
	}
	return ret
}

func main() {
	var configFlag = flag.String("config", "", "")
	flag.Parse()

	var config Config

	config_serialized, err := ioutil.ReadFile(*configFlag)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(config_serialized, &config)
	if err != nil {
		log.Panic(err)
	}
	results := make([][]pf.Security, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	get :=
		func(out_index int, s pf.SecurityFetcher) {
			prices, err := s.Run()
			if err != nil {
				log.Panic(err)
			}
			results[out_index] = prices
			wg.Done()
		}
	go get(0, config.V)
	go get(1, config.X)
	wg.Wait()

	out := []pf.Security{}
	for _, r := range results {
		out = append(out, r...)
	}
	namePadWidth := longestName(out)
	pricePadWidth := longestPrice(out)
	for _, x := range out {
		currencyFmt := fmt.Sprintf("$%.02f", x.Price)
		fmt.Printf("P %s %-*s %*s\n", time.Now().Format("2006/01/02"), namePadWidth, x.Name, pricePadWidth, currencyFmt)
	}
}
