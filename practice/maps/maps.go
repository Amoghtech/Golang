package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google": "https://google.com",
	// 	"Amazon Web Services": "https://aws.com",
	// }
	// websites["LinkedIn"] = "https://linkedin.com"
	// fmt.Println(websites)
	// delete(websites, "Google")
	// fmt.Println(websites)

	userNames := make([]string, 2, 5)
	// userNames := [5]string{"Hello"}
	userNames[0] = "Hello"
	userNames[1] = "Hello"
	userNames[2] = "Hello"
	fmt.Println(userNames)

	// coursesRatings := make(floatMap, 3)

	// coursesRatings["go"] = 4.7
	// coursesRatings["react"] = 4.8
	// coursesRatings["node"] = 4.9
	// coursesRatings.output()

	// for index,value := range coursesRatings {
	// 	fmt.Println(index, value)
	// }
}
