package main

import (
	"fmt"
)

func main() {
	// var productNames [2]string
	// productNames = [2]string{"Name", "Hello"}
	// fmt.Println(productNames[1])

	// prices := [10]float64{10.99, 12.22, 43, 22, 67, 44, 67, 12, 90, 92}
	// fmt.Println(prices[2])

	// featuredPrices := prices[2:]
	// highLightedPrice := featuredPrices[:3]
	// fmt.Println(len(highLightedPrice), cap(highLightedPrice))

	// highLightedPrice = highLightedPrice[:8]
	// fmt.Println(len(highLightedPrice), cap(highLightedPrice))

	prices := []float64{10, 99, 20}
	prices[1] = 9.99
	fmt.Println(prices)
	prices = append(prices, 5.99)
	fmt.Println(prices)

}
