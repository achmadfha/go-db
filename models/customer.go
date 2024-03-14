package models

type Customer struct {
	CustomerID   int    `json:"customer_id"`
	NamaCustomer string `json:"nama_customer"`
	NoHP         string `json:"no_hp"`
}
