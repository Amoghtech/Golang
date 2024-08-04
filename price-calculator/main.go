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
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make(chan error)
	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		job := priceJob.NewTaxIncludedPriceJob(fm, taxRate)
		doneChans[index] = make(chan bool)
		// errorChans[index] = make(chan error)
		go job.Process(doneChans[index], errorChans)
		// <- doneChans[index]

		// if(err!=nil){
		// 	fmt.Println("Could not process job", err)
		// 	return
		// }
	}

	err := <- errorChans
	fmt.Println(err)
	// for index := range taxRates {
	// 	select {
	// 		case err := <-errorChans[index]:
	// 			if(err!=nil){
	// 				fmt.Println(err)
	// 			}
	// 		case <-doneChans[index]:
	// 			fmt.Println("Done!")
	// 	}
	// }

	// for _, error := range errorChans {
	// 	<-error
	// }

	// for _, done := range doneChans {
	// 	<-done
	// }
}
