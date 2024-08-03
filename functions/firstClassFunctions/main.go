package main

import (
	"fmt"
)

type TransformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	// doubled := transformNumbers(&numbers, double)
	doubled := transformNumbers(&numbers, getTransformerFunction(&numbers))
	tripled := transformNumbers(&moreNumbers, getTransformerFunction(&moreNumbers))
	fmt.Println(doubled, tripled)
}

func transformNumbers(numbers *[]int, transform TransformFunc) []int {
	dNumbers := make([]int, len(*numbers))
	for index, value := range *numbers {
		dNumbers[index] = transform(value)
	}
	return dNumbers
}

func getTransformerFunction(number *[]int) TransformFunc {
	if (*number)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
