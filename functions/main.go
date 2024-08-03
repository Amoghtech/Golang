package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(sumup(numbers...))
}

func sumup(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
