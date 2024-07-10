package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	for {
		fmt.Println(`What services you want:
		1. Check Balance
		2. Deposit Money
		3. Withdraw Money
		4. Exit
		`)

		var choice int
		var accountBalance, err = fileops.GetBalanceFromFile(accountBalanceFile)

		if err != nil {
			fmt.Println("ERROR!")
			fmt.Print(err)
			fmt.Println("------------------------------")
			return
		}

		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")
				continue
			}
			fmt.Println("New Balance:", accountBalance)
			accountBalance += depositAmount
			fileops.WriteBalance(accountBalance, accountBalanceFile)
		case 3:
			var withdrawMoney float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawMoney)
			accountBalance -= withdrawMoney
			fmt.Println("New Balance:", accountBalance)
			fileops.WriteBalance(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}
