package main

import (
	"fmt"
	"os"
)

func storeResults(ebt ,profit, ratio float64){
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), os.ModeAppend)
}

func main() {
	var revenue, expenses, taxRate float64

	outputText("Revenue: ", &revenue)
	outputText("Expenses: ", &expenses)
	outputText("Tax Rate: ", &taxRate)

	fmt.Print(revenue, expenses, taxRate)

	ebt, profit, ratio := calculateFutureValues(revenue, expenses, taxRate)



	formattedEBT := fmt.Sprintf("EBT: %.2f\n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %.2f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedProfit, formattedRatio)

	storeResults(ebt, profit, taxRate)
}

func outputText(outputText string, outputVariable *float64) {
	fmt.Print(outputText)
	fmt.Scan(outputVariable)
	if(*outputVariable<=0){
		panic("Invalid " + outputText + ", must be greater bthan 0");
	}
}

func calculateFutureValues(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = (revenue - expenses)
	profit = ((revenue - expenses) * (1 - taxRate/100))
	ratio = ebt/profit
	return
}
