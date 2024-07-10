package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetBalanceFromFile(file string) (float64, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return 1020, errors.New("error to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1030, errors.New("error to parse float")
	}

	return balance, nil
}

func WriteBalance(balance float64, file string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(file, []byte(balanceText), 0644)
}
