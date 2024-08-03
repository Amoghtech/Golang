package practice

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"BasketBall", "Football", "Cricket"}
	fmt.Println("Hobbies are: ", hobbies)

	fmt.Println("First Element: ", hobbies[0])
	combinedHobby := hobbies[1:]
	fmt.Println("Combined List: ", combinedHobby)

	newSlice1 := hobbies[:2]
	// newSlice:= hobbies[0:2]
	fmt.Println("NewSlice(First and Second): ", newSlice1)

	newSlice2 := newSlice1[1:3]
	fmt.Println("NewSlice(Second and Third): ", newSlice2)

	newSlice2 = append(newSlice2, newSlice1...)

	goals := []string{"Cricketer", "Footballer"}
	goals[1] = "Teacher"
	goals = append(goals, "Doctor")
	fmt.Println("Goals are: ", goals)

	products := []Product{Product{
		"Samosa",
		1,
		10,
	}, Product{
		"Momos",
		2,
		20,
	}}

	products = append(products, Product{
		"Pizza",
		3,
		90,
	})
	fmt.Println("Products are: ", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
