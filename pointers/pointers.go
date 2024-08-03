package main

import (
	"fmt"
)

func main() {
	age := 32
	agePointer := &age
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age -= 18 
}


