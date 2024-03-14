package services

import (
	"challenge-godb/models"
	"challenge-godb/repositories"
	"log"
)

func GetAllCustomers() ([]models.Customer, error) {
	return repositories.GetAllCustomers()
}

func ViewCustomerByID(customerID int) ([]models.Customer, error) {

	customer, err := repositories.ViewCustomerByID(customerID)
	if err != nil {
		log.Println("Error retrieving customer:", err)
		return nil, err
	}
	return customer, nil
}

func AddCustomer(customer models.Customer) error {

	err := repositories.AddCustomer(customer)
	if err != nil {
		log.Println("Error adding customer:", err)
	}
	return nil
}

func UpdateCustomerByID(customerID int, updatedCustomer models.Customer) error {

	err := repositories.UpdateCustomerByID(customerID, updatedCustomer)
	if err != nil {
		log.Println("Error updating customer in service:", err)
		return err
	}
	return nil
}

func DeleteCustomer(customerID int) error {

	err := repositories.DeleteCustomer(customerID)
	if err != nil {
		log.Println("Error Delete in Service :", err)
	}

	return nil
}

func GetCustomerByPhone(NoHP string) ([]models.Customer, error) {

	customer, err := repositories.GetCustomerByPhone(NoHP)
	if err != nil {
		log.Println("Error retrieving customer:", err)
	}

	return customer, nil
}
