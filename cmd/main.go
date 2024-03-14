package main

import (
	"challenge-godb/cmd/submenu"
	"challenge-godb/db"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"log"
)

func main() {

	err := db.GetDB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	myFigure := figure.NewFigure("Loundry Apps", "standard", true)
	myFigure.Print()
	CommandMenuLoop()
	defer db.CloseDB()
}

func CommandMenuLoop() {
	for {
		fmt.Println("Main Menu\n")
		fmt.Println("1. Customer")
		fmt.Println("2. Services")
		fmt.Println("3. Transactions")
		fmt.Println("4. Exit \n")
		fmt.Println("\n")

		fmt.Print("Please enter the number of your choice (1-4): ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			submenu.CustomerSubMenu()
		case 2:
			submenu.PelayananSubMenu()
		case 3:
			submenu.TransactionsSubMenu()
		case 4:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter a number between 1 and 4.\n")
		}
	}
}
