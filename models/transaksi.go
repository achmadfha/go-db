package models

import (
	"database/sql"
	"time"
)

type TransaksiLaundry struct {
	TransaksiID    int           `json:"transaksi_id"`
	NoNota         string        `json:"no_nota"`
	TanggalMasuk   time.Time     `json:"tanggal_masuk"`
	TanggalSelesai time.Time     `json:"tanggal_selesai"`
	DiterimaOleh   string        `json:"diterima_oleh"`
	CustomerID     sql.NullInt64 `json:"customer_id"`
}
