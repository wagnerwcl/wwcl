package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.8
	courseRating["react"] = 4.9
	courseRating["angular"] = 4.7

	courseRating.output()

	for key, value := range userNames {
		fmt.Println(key)
		fmt.Println(value)
	}
}
