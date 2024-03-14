package services

import (
	"challenge-godb/models"
	"challenge-godb/repositories"
	"log"
)

func GetAllServices() ([]models.Pelayanan, error) {
	return repositories.GetAllServices()
}

func AddNewServices(pelayanan models.Pelayanan) error {

	err := repositories.AddNewServices(pelayanan)
	if err != nil {
		log.Println("Error Add New Services :", err)
		return err
	}
	return nil
}

func ViewServicesByID(pelayananID int) ([]models.Pelayanan, error) {

	pelayanan, err := repositories.ViewServicesByID(pelayananID)
	if err != nil {
		log.Println("Error retrieving services:", err)
		return nil, err
	}
	return pelayanan, nil
}

func UpdateServicesByID(pelayananID int, updatedServices models.Pelayanan) error {

	err := repositories.UpdateServicesByID(pelayananID, updatedServices)
	if err != nil {
		log.Println("Error updating Service on service:", err)
		return err
	}
	return nil
}
