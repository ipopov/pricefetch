package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	pf "github.com/ipopov/pricefetch/lib"
)

type output struct {
	name string
	// TODO: Convert this to decimal.
	price float64
}

func longestName(xs []output) int {
	ret := 0
	for _, x := range xs {
		if len(x.name) > ret {
			ret = len(x.name)
		}
	}
	return ret
}

func longestPrice(xs []output) int {
	ret := 0
	for _, x := range xs {
		// 4 is for the dollar sign, dot, cents.
		digits := 4 + 1 + int(math.Log10(x.price))
		if digits > ret {
			ret = digits
		}
	}
	return ret
}

func main() {
	out := make([]output, len(config))
	var wg sync.WaitGroup
	wg.Add(len(config))
	for i, v := range config {
		go func(out_index int, s pf.Security) {
			price, err := s.GetPrice()
			if err != nil {
				log.Panic(err)
			}
			out[out_index] = output{s.GetName(), price}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	namePadWidth := longestName(out)
	pricePadWidth := longestPrice(out)
	for _, x := range out {
		currencyFmt := fmt.Sprintf("$%.02f", x.price)
		fmt.Printf("P %s %-*s %*s\n", time.Now().Format("2006/01/02"), namePadWidth, x.name, pricePadWidth, currencyFmt)
	}
}
