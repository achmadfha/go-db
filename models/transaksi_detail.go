package models

type TransaksiDetail struct {
	TransaksiDetailID int `json:"transaksi_detail_id"`
	TransaksiID       int `json:"transaksi_id"`
	PelayananID       int `json:"pelayanan_id"`
	Jumlah            int `json:"jumlah"`
}
