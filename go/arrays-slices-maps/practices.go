package main

import "fmt"

func main() {
	hobbies := [3]string{"reading", "coding", "gaming"}
	// 1)
	fmt.Println(hobbies)
	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	// 3)
	newList := hobbies[:2]
	fmt.Println(newList)
	// 4)
	slice4 := hobbies[1:3]
	fmt.Println(slice4)
	// 5)
	dynamicArray := []string{"Learn Go", "Learn Python"}
	fmt.Println(dynamicArray)
	// 6)
	dynamicArray[1] = "Learn AI"
	dynamicArray = append(dynamicArray, "Learn QA")
	fmt.Println(dynamicArray)
	// 7)
	type product struct {
		title string
		id    int
		price float64
	}

	products := []product{
		{"Book", 1, 10.99},
		{"Laptop", 2, 999.99},
	}

	fmt.Println(products)
	// 8)
	products = append(products, product{"Phone", 3, 599.99})
	fmt.Println(products)
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
