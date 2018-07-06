package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	pf "github.com/ipopov/pricefetch/lib"
)

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
	results := make([][]pf.Security, len(config))
	var wg sync.WaitGroup
	wg.Add(len(config))
	for i, v := range config {
		go func(out_index int, s pf.SecurityFetcher) {
			prices, err := s.Run()
			if err != nil {
				log.Panic(err)
			}
			results[out_index] = prices
			wg.Done()
		}(i, v)
	}
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
