CREATE TABLE Customer
(
    customer_id   SERIAL PRIMARY KEY,
    nama_customer VARCHAR(255) NOT NULL,
    no_hp         VARCHAR(15)  NOT NULL
);

CREATE TABLE Pelayanan
(
    pelayanan_id    SERIAL PRIMARY KEY,
    jenis_pelayanan VARCHAR(50) NOT NULL,
    satuan          VARCHAR(10) NOT NULL,
    harga           INTEGER     NOT NULL
);

CREATE TABLE Transaksi_Laundry
(
    transaksi_id     SERIAL PRIMARY KEY,
    no_nota          VARCHAR(10) NOT NULL,
    tanggal_masuk    DATE        NOT NULL,
    tanggal_selesai  DATE        NOT NULL,
    diterima_oleh    VARCHAR(50) NOT NULL,
    customer_id      INTEGER REFERENCES Customer (customer_id) ON DELETE SET NULL
);

CREATE TABLE Transaksi_Detail
(
    transaksi_detail_id SERIAL PRIMARY KEY,
    transaksi_id        INTEGER REFERENCES Transaksi_Laundry (transaksi_id) ON DELETE SET NULL,
    pelayanan_id        INTEGER REFERENCES Pelayanan (pelayanan_id) ON DELETE SET NULL,
    jumlah              INTEGER NOT NULL
);
