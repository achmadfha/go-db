package services

import (
	"challenge-godb/models"
	"challenge-godb/repositories"
)

func GetAllTransactions() ([]models.TransaksiLaundry, error) {
	return repositories.GetAllTransactions()
}
