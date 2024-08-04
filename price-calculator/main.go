package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	priceJob "example.com/price-calculator/prices"
)


func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		job := priceJob.NewTaxIncludedPriceJob(fm, taxRate)
		err := job.Process()
		if(err!=nil){
			fmt.Println("Could not process job", err)
			return
		}
	}
}
