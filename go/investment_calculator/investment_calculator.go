package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	output("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	output("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	output("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.2f\nReal Value: %.2f\n", futureValue, futureRealValue)

	fmt.Print(``)
}

func output(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
