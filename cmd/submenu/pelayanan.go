package submenu

import (
	"challenge-godb/handlers"
	"fmt"
)

func PelayananSubMenu() {
	for {
		fmt.Println("\nServices Menu: \n")
		fmt.Println("1. View All Services")
		fmt.Println("2. Add New Services")
		fmt.Println("3. Update Services")
		fmt.Println("4. Back to Main Menu")
		fmt.Println("\n")

		fmt.Print("Please enter the number of your choice (1-4):")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.ViewAllServices()
		case 2:
			handlers.AddServices()
		case 3:
			handlers.UpdateSevices()
		case 4:
			fmt.Println("Return Main Menu")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter a number between 1 and 4.\n")
		}
	}
}
