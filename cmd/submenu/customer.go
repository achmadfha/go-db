package submenu

import (
	"challenge-godb/handlers"
	"fmt"
)

func CustomerSubMenu() {
	for {
		fmt.Println("\nCustomer Menu: \n")
		fmt.Println("1. View All Customers")
		fmt.Println("2. Add New Customer")
		fmt.Println("3. Update Customer Info")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. Back to Main Menu")
		fmt.Println("\n")

		fmt.Print("Please enter the number of your choice (1-5):")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.ViewCustomers()
		case 2:
			handlers.AddCustomer()
		case 3:
			handlers.UpdateCustomerByID()
		case 4:
			handlers.DeleteCustomerByID()
		case 5:
			fmt.Println("Return Main Menu")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter a number between 1 and 5.\n")
		}
	}
}
