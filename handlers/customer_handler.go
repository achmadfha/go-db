package handlers

import (
	"bufio"
	"challenge-godb/models"
	"challenge-godb/services"
	"challenge-godb/utils"
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"os"
	"strings"
)

func ViewCustomers() {
	customers, err := services.GetAllCustomers()
	if err != nil {
		log.Fatal(err)
	}

	if len(customers) == 0 {
		fmt.Println("No Customer Found Sorry :(")
		return
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "No"},
			{Text: "Name"},
			{Text: "Phone"},
		},
	}

	for _, customer := range customers {
		row := []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", customer.CustomerID)},
			{Text: customer.NamaCustomer},
			{Text: customer.NoHP},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	fmt.Println("\nAll Customer :")
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func AddCustomer() {

	fmt.Print("\nEnter customer full name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	customerName := scanner.Text()

	var customerPhone string
	for {
		fmt.Print("Customer phone number (12 digits max): ")
		fmt.Scanln(&customerPhone)

		if utils.IsValidPhoneNumber(customerPhone) {
			break
		}
		fmt.Println("Invalid phone number ! Please enter a valid 12-digit phone number.")
	}

	newCustomer := models.Customer{
		NamaCustomer: customerName,
		NoHP:         customerPhone,
	}

	fmt.Print("Are you sure you want to add this customer? (y/n): ")

	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		services.AddCustomer(newCustomer)
		fmt.Println("New customer added successfully!\n")
	} else {
		fmt.Println("Customer addition canceled.\n")
	}
}

func UpdateCustomerByID() {
	fmt.Print("\nEnter customer ID: ")
	var customerID int
	fmt.Scanln(&customerID)

	// get old data
	existingCustomer, err := services.ViewCustomerByID(customerID)
	if err != nil {
		log.Println("Error checking customer existence:", err)
		return
	}

	// customer not found
	if len(existingCustomer) == 0 {
		fmt.Println("Customer with ID", customerID, "not found.")
		return
	}

	existingCustomers := existingCustomer
	updatedCustomer := readCustomerDetails(&existingCustomers[0])

	fmt.Print("Are you sure you want to update this customer? (y/n): ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		err := services.UpdateCustomerByID(customerID, updatedCustomer)
		if err != nil {
			log.Println("Error updating customer:", err)
			return
		}
		fmt.Println("Customer updated successfully!")
	} else {
		fmt.Println("Customer update canceled.")
	}

}

func readCustomerDetails(existingCustomer *models.Customer) models.Customer {
	var updatedCustomer models.Customer

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Updated Customer Name (%s): ", existingCustomer.NamaCustomer)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		updatedCustomer.NamaCustomer = existingCustomer.NamaCustomer
	} else {
		updatedCustomer.NamaCustomer = name
	}

	for {
		fmt.Printf("Enter Updated Phone Number (%s): ", existingCustomer.NoHP)
		phoneNumber, _ := reader.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)

		if phoneNumber == "" {
			updatedCustomer.NoHP = existingCustomer.NoHP
			break
		}

		if utils.IsValidPhoneNumber(phoneNumber) {
			updatedCustomer.NoHP = phoneNumber
			break
		} else {
			fmt.Println("Invalid phone number. Please enter a valid 13-digit phone number.")
		}
	}

	return updatedCustomer
}

func DeleteCustomerByID() {

	fmt.Print("\nEnter customer ID: ")
	var customerID int
	fmt.Scanln(&customerID)

	err := services.DeleteCustomer(customerID)
	if err != nil {
		log.Println("Error deleting customer:", err)
		return
	}

	fmt.Printf("Customer with ID %d deleted successfully\n", customerID)
}
