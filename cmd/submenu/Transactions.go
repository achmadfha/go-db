package submenu

import (
	"challenge-godb/handlers"
	"fmt"
)

func TransactionsSubMenu() {
	for {
		fmt.Println("\nTransactions Menu: \n")
		fmt.Println("1. View Transactions ")
		fmt.Println("2. Add Transactions")
		fmt.Println("3. Back to Main Menu")
		fmt.Println("\n")

		fmt.Print("Please enter the number of your choice (1-3):")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.ViewAllTransactions()
		case 2:
			handlers.AddTransactions()
		case 3:
			fmt.Println("Return Main Menu")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter a number between 1 and 3.\n")
		}
	}
}
