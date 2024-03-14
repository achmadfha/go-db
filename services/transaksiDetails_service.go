package services

import (
	"challenge-godb/models"
	"challenge-godb/repositories"
	"log"
)

func AddTransactionDetails(transaksidetail models.TransaksiDetail) error {

	err := repositories.AddTransactionDetails(transaksidetail)
	if err != nil {
		log.Println("Error adding transaksidetail: ", err)
	}
	return nil
}
