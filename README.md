## Getting Started

### Database Setup

*  Buka challenge-godb/db/migrations
*  Buka pgAdmin lalu connect pada PostgreSQL server.
*  Create new database
*  Click Kanan pada database yang telah dibuat lalu pilih "Query Tool,"
*  Pilih Open file > ddl.sql lalu Execute 
*  Buka git bash pada folder project
*  Copy file .env
```bash
cp .env.example .env
```
*  Set up .env 
*  Kembali pada pgAdmin
*  Pada "Query Tool," Pilih open file > dml.sql lalu Execute


### Run Application

Buka git bash / cmd pada folder project challenge-godb
lalu tinggal jalankan app dengan perintah berikut
```bash
go run cmd/main.go
```

### Notes
Pastikan .env sudah sesuai dengan koneksi db kalian