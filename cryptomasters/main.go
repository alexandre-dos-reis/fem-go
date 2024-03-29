package main

import (
	"fem-go/cryptomasters/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(c string) {
			getCurrencyData(c)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Printf("An error occured %v\n", err)
	}
	fmt.Printf("The rate for %v is %2.f\n", rate.Currency, rate.Price)
}
