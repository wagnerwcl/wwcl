package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//	=> Show error message & exit if invalid input is provided
//	- No negative numbers
//  - Not 0
// 2) Store calculated result into file

const valueFile = "values.txt"

func main() {
	revenue, err := userData("Revenue:")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := userData("Expenses:")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := userData("Tax:")

	if err != nil {
		fmt.Println(err)
		return
	}

	calcValues(revenue, expenses, taxRate)
}

func userData(insertText string) (float64, error) {
	var userInput float64
	fmt.Print(insertText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("invalid value")
	}

	return userInput, nil
}

func calcValues(revenue, expenses, taxRate float64) {

	ebt := revenue - expenses
	profit := ebt - taxRate
	ratio := ebt / profit

	writeValue(ebt, profit, ratio)
}

func writeValue(ebt, profit, ratio float64) {
	valueText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile(valueFile, []byte(valueText), 0644)
}
