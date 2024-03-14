package models

type Pelayanan struct {
	PelayananID    int    `json:"pelayanan_id"`
	JenisPelayanan string `json:"jenis_pelayanan"`
	Satuan         string `json:"satuan"`
	Harga          int    `json:"harga"`
}
