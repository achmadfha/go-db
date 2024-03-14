package repositories

import (
	"challenge-godb/db"
	"challenge-godb/models"
	"log"
)

func GetAllCustomers() ([]models.Customer, error) {

	query := "SELECT * FROM Customer"
	rows, err := db.GetDB().Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.CustomerID, &customer.NamaCustomer, &customer.NoHP)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func ViewCustomerByID(customerID int) ([]models.Customer, error) {

	query := "SELECT * FROM Customer WHERE customer_id = $1"
	rows, err := db.GetDB().Query(query, customerID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.CustomerID, &customer.NamaCustomer, &customer.NoHP)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func AddCustomer(customer models.Customer) error {

	query := "INSERT INTO Customer (nama_customer, no_hp) VALUES ($1, $2) RETURNING customer_id"
	err := db.GetDB().QueryRow(query, customer.NamaCustomer, customer.NoHP).Scan(&customer.CustomerID)
	if err != nil {
		log.Println("Error inserting customer:", err)
		return err
	}

	return nil
}

func UpdateCustomerByID(customerID int, updatedCustomer models.Customer) error {

	query := `UPDATE Customer SET nama_customer = $1, no_hp = $2 WHERE customer_id = $3`
	_, err := db.GetDB().Exec(query, updatedCustomer.NamaCustomer, updatedCustomer.NoHP, customerID)
	if err != nil {
		log.Println("Error updating customer:", err)
	}
	return nil
}

func DeleteCustomer(customerID int) error {

	query := "DELETE FROM Customer WHERE customer_id = $1"
	_, err := db.GetDB().Exec(query, customerID)
	if err != nil {
		log.Println("Error delete customer :", err)
	}

	return nil
}

func GetCustomerByPhone(NoHP string) ([]models.Customer, error) {

	query := "SELECT * FROM Customer WHERE no_hp = $1"
	rows, err := db.GetDB().Query(query, NoHP)
	if err != nil {
		log.Println("Error Executing query:", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.CustomerID, &customer.NamaCustomer, &customer.NoHP)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
