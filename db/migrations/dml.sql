INSERT INTO Customer (nama_customer, no_hp)
VALUES ('Jessica', '0812338947854'),
       ('John Doe', '0821123456789'),
       ('Alice Johnson', '0815777888999'),
       ('Bob Smith', '0831234567890'),
       ('Eva Martinez', '0812345678901'),
       ('Michael Williams', '0823456789012'),
       ('Sophia Davis', '0834567890123'),
       ('Daniel Wilson', '0815678901234'),
       ('Olivia Brown', '0826789012345'),
       ('Mason Miller', '0837890123456');

INSERT INTO Pelayanan (jenis_pelayanan, satuan, harga)
VALUES ('Cuci + Setrika', 'KG', 7000),
       ('Laundry Bed Cover', 'buah', 50000),
       ('Laundry Boneka', 'buah', 250000),
       ('Express Service', 'buah', 100000),
       ('Dry Cleaning', 'KG', 12000),
       ('Fold Only', 'buah', 3000),
       ('Sneaker Cleaning', 'buah', 15000),
       ('Cuci + Setrika Express', 'KG', 10000),
       ('Bed Linen Laundry', 'buah', 60000),
       ('Stuffed Animal Cleaning', 'buah', 20000);

INSERT INTO Transaksi_Laundry (no_nota, tanggal_masuk, tanggal_selesai, diterima_oleh, customer_id)
VALUES ('12345', '2022-08-18', '2022-08-20', 'Mirna', 1),
       ('67890', '2022-09-10', '2022-09-12', 'Andy', 2),
       ('54321', '2022-08-25', '2022-08-27', 'Lisa', 3),
       ('98765', '2022-09-15', '2022-09-17', 'David', 4),
       ('45678', '2022-09-05', '2022-09-08', 'Sophie', 5),
       ('23456', '2022-09-20', '2022-09-22', 'George', 6),
       ('78901', '2022-09-08', '2022-09-10', 'Emma', 7),
       ('34567', '2022-09-03', '2022-09-06', 'James', 8),
       ('89012', '2022-09-12', '2022-09-15', 'Ava', 9),
       ('12378', '2022-09-18', '2022-09-20', 'Noah', 10);

INSERT INTO Transaksi_Detail (transaksi_id, pelayanan_id, jumlah)
VALUES (1, 1, 5),
       (1, 2, 1),
       (1, 3, 2),
       (2, 4, 3),
       (2, 5, 2),
       (3, 2, 1),
       (3, 3, 4),
       (4, 6, 2),
       (4, 7, 1),
       (5, 8, 3),
       (5, 9, 1),
       (6, 10, 2),
       (6, 1, 4),
       (7, 2, 3),
       (7, 3, 1),
       (8, 4, 2),
       (8, 5, 1),
       (9, 6, 3),
       (9, 7, 2),
       (10, 8, 1),
       (10, 9, 2);
