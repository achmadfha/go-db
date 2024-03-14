package repositories

import (
	"challenge-godb/db"
	"challenge-godb/models"
	"log"
)

func AddTransactionDetails(transaksidetail models.TransaksiDetail) error {

	query := "INSERT INTO Transaksi_Detail (transaksi_id, pelayanan_id, jumlah) VALUES ($1, $2, $3) RETURNING transaksi_detail_id"
	row := db.GetDB().QueryRow(query, transaksidetail.TransaksiID, transaksidetail.PelayananID, transaksidetail.Jumlah)

	var transaksiDetailID int
	if err := row.Scan(&transaksiDetailID); err != nil {
		log.Println("Error scanning Transaction Details: ", err)
		return err
	}

	return nil
}
