package repositories

import (
	"challenge-godb/db"
	"challenge-godb/models"
	"log"
)

func GetAllServices() ([]models.Pelayanan, error) {

	query := "SELECT * FROM Pelayanan"
	rows, err := db.GetDB().Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var pelayanans []models.Pelayanan
	for rows.Next() {
		var pelayanan models.Pelayanan
		err := rows.Scan(&pelayanan.PelayananID, &pelayanan.JenisPelayanan, &pelayanan.Satuan, &pelayanan.Harga)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		pelayanans = append(pelayanans, pelayanan)
	}

	return pelayanans, nil
}

func AddNewServices(pelayanan models.Pelayanan) error {

	query := "INSERT INTO Pelayanan (jenis_pelayanan, satuan, harga) VALUES ($1, $2, $3) RETURNING pelayanan_id"
	err := db.GetDB().QueryRow(query, pelayanan.JenisPelayanan, pelayanan.Satuan, pelayanan.Harga).Scan(&pelayanan.PelayananID)
	if err != nil {
		log.Println("Error Inserting customer :", err)
		return nil
	}

	return nil
}

func ViewServicesByID(pelayananID int) ([]models.Pelayanan, error) {

	query := "SELECT * FROM Pelayanan WHERE pelayanan_id = $1"
	rows, err := db.GetDB().Query(query, pelayananID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var pelayanan []models.Pelayanan
	for rows.Next() {
		var p models.Pelayanan
		err := rows.Scan(&p.PelayananID, &p.JenisPelayanan, &p.Satuan, &p.Harga)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		pelayanan = append(pelayanan, p)
	}
	return pelayanan, nil
}

func UpdateServicesByID(PelayananID int, updatedServices models.Pelayanan) error {

	query := `UPDATE Pelayanan SET jenis_pelayanan = $1, satuan = $2, harga = $3 WHERE pelayanan_id = $4`
	_, err := db.GetDB().Exec(query, updatedServices.JenisPelayanan, updatedServices.Satuan, updatedServices.Harga, PelayananID)
	if err != nil {
		log.Println("Error Updating Services:", err)
		return err
	}
	return nil
}

func CreateNewServices(pelayanan models.Pelayanan) (int, error) {

	query := "INSERT INTO Pelayanan (jenis_pelayanan, satuan, harga) VALUES ($1, $2, $3) RETURNING pelayanan_id"
	err := db.GetDB().QueryRow(query, pelayanan.JenisPelayanan, pelayanan.Satuan, pelayanan.Harga).Scan(&pelayanan.PelayananID)
	if err != nil {
		log.Println("Error Inserting customer :", err)
		return 0, err
	}

	return pelayanan.PelayananID, nil
}
