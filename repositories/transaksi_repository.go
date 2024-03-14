package repositories

import (
	"challenge-godb/db"
	"challenge-godb/models"
	"log"
)

func GetAllTransactions() ([]models.TransaksiLaundry, error) {
	query := "SELECT * FROM Transaksi_Laundry"
	rows, err := db.GetDB().Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []models.TransaksiLaundry

	for rows.Next() {
		var transaction models.TransaksiLaundry
		err := rows.Scan(
			&transaction.TransaksiID,
			&transaction.NoNota,
			&transaction.TanggalMasuk,
			&transaction.TanggalSelesai,
			&transaction.DiterimaOleh,
			&transaction.CustomerID,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error reading rows:", err)
		return nil, err
	}

	return transactions, nil
}

func AddTransactions(transaksiloundry models.TransaksiLaundry) (int, error) {

	query := "INSERT INTO Transaksi_Laundry (no_nota, tanggal_masuk, tanggal_selesai, diterima_oleh, customer_id) VALUES ($1, $2, $3, $4, $5) RETURNING transaksi_id"
	err := db.GetDB().QueryRow(query, transaksiloundry.NoNota, transaksiloundry.TanggalMasuk, transaksiloundry.TanggalSelesai, transaksiloundry.DiterimaOleh, transaksiloundry.CustomerID).Scan(&transaksiloundry.TransaksiID)
	if err != nil {
		log.Println("Error inserting transaksi loundry: ", err)
		return 0, err
	}

	return transaksiloundry.TransaksiID, nil
}
