package main

import "fmt"

func main() {
	prices := []float64{10.99, 9.99}
	prices = append(prices, 5.99)
	fmt.Println(prices)

	discountPrices := []float64{100.00, 80.99, 20, 59}
	prices = append(prices, discountPrices...) // Join lists
	fmt.Println(prices)
}

// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices[2])
// 	prices[3] = 35.5
// 	fmt.Println(prices[3])

// 	featurePrices := prices[1:3] //Slice
// 	highletedPrices := featurePrices[1:]

// 	featurePrices[0] = 100.0

// 	fmt.Println(featurePrices)
// 	fmt.Println(highletedPrices)
// 	fmt.Println(len(featurePrices), cap(featurePrices))
// }
