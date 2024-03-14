package handlers

import (
	"bufio"
	"challenge-godb/models"
	"challenge-godb/repositories"
	"challenge-godb/services"
	"challenge-godb/utils"
	"database/sql"
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func ViewAllTransactions() {
	transactions, err := services.GetAllTransactions()
	if err != nil {
		log.Fatal(err)
	}

	if len(transactions) == 0 {
		fmt.Println("No Transactions found sorry :(")
	}

	tabel := simpletable.New()
	tabel.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "No"},
			{Text: "No Nota"},
			{Text: "Tanggal Masuk"},
			{Text: "Tanggal Selesai"},
			{Text: "Diterimah Oleh"},
			{Text: "Customer Name"},
			{Text: "NoHP"},
		},
	}

	for i, transaction := range transactions {

		tanggalMasuk := transaction.TanggalMasuk.Format("02-01-2006")
		tanggalSelesai := transaction.TanggalSelesai.Format("02-01-2006")

		var customerID int
		if transaction.CustomerID.Valid {
			customerID = int(transaction.CustomerID.Int64)
		} else {
			continue
		}

		customer, err := services.ViewCustomerByID(customerID)
		if err != nil {
			log.Println("Error getting customer information:", err)
			return
		}

		if len(customer) > 0 {
			customer := customer[0]

			row := []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", i+1)},
				{Text: transaction.NoNota},
				{Text: tanggalMasuk},
				{Text: tanggalSelesai},
				{Text: transaction.DiterimaOleh},
				{Text: customer.NamaCustomer},
				{Text: customer.NoHP},
			}

			tabel.Body.Cells = append(tabel.Body.Cells, row)
		}
	}

	fmt.Println("\nAll Transaction")
	tabel.SetStyle(simpletable.StyleUnicode)
	fmt.Println(tabel.String())
}

func AddTransactions() {

	var customerPhone string
	for {
		fmt.Print("\nEnter Customer Phone Number: ")
		fmt.Scanln(&customerPhone)

		if utils.IsValidPhoneNumber(customerPhone) {
			break
		}
		fmt.Println("Invalid phone number ! Please enter a valid 12-digit phone number.")
	}

	existingCustomer, err := services.GetCustomerByPhone(customerPhone)
	if err != nil {
		log.Println("Error checking customer existence:", err)
		return
	}

	if len(existingCustomer) == 0 {
		fmt.Println("Customer with phone number", customerPhone, "not found.")
		return
	}

	firstExistingCustomer := existingCustomer[0]
	customerID := firstExistingCustomer.CustomerID

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(90000) + 10000
	noNota := strconv.Itoa(randomNumber)

	DateIn, err := utils.ReadInput("Tanggal Masuk (dd-mm-yyyy): ", utils.ValidateDate)
	if err != nil {
		fmt.Println("Error date:", err)
		return
	}
	parsedDateIn, ok := DateIn.(time.Time)
	if !ok {
		fmt.Println("Error: DateIn is not of type time.Time")
		return
	}

	DateOut, err := utils.ReadInput("Tanggal Selesai (dd-mm-yyyy): ", utils.ValidateDate)
	if err != nil {
		fmt.Println("Error date:", err)
		return
	}

	parsedDateOut, ok := DateOut.(time.Time)
	if !ok {
		fmt.Println("Error: DateIn is not of type time.Time")
		return
	}

	fmt.Print("Diterima Oleh: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	diterimaOleh := scanner.Text()

	fmt.Print("Enter Pelayanan name: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	servicesName := scanner.Text()

	pelayananJumlah := utils.ValidationJumlah()

	var price string
	for {
		fmt.Print("Harga: ")
		fmt.Scanln(&price)

		if utils.IsValidNumber(price) {
			break
		}
		fmt.Println("Invalid Price ! Please enter a valid number")
	}

	prices, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("Error converting price to integer:", err)
		return
	}

	var jumlahString string
	for {
		fmt.Print("Jumlah: ")
		fmt.Scanln(&jumlahString)

		if utils.IsValidNumber(jumlahString) {
			break
		}
		fmt.Println("Invalid Price! Please enter a valid number")
	}

	jumlah, err := strconv.Atoi(jumlahString)
	if err != nil {
		fmt.Println("Error converting price to integer:", err)
		return
	}

	fmt.Print("Are you sure you want to add this transactions ? (y/n): ")

	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {

		addPelayanan := models.Pelayanan{
			JenisPelayanan: servicesName,
			Satuan:         pelayananJumlah,
			Harga:          prices,
		}

		trxID, err := repositories.CreateNewServices(addPelayanan)
		if err != nil {
			log.Println("Error adding pelayanan:", err)
			return
		}

		addTransactions := models.TransaksiLaundry{
			NoNota:         noNota,
			TanggalMasuk:   parsedDateIn,
			TanggalSelesai: parsedDateOut,
			DiterimaOleh:   diterimaOleh,
			CustomerID:     sql.NullInt64{Int64: int64(customerID), Valid: true},
		}
		trxDetailsID, err := repositories.AddTransactions(addTransactions)
		if err != nil {
			log.Println("Error adding transaction:", err)
			return
		}

		addTransactionsDetails := models.TransaksiDetail{
			TransaksiID: trxDetailsID,
			PelayananID: trxID,
			Jumlah:      jumlah,
		}
		services.AddTransactionDetails(addTransactionsDetails)
		fmt.Println("New Services added successfully!\n")
	} else {
		fmt.Println("Customer addition canceled.\n")
	}

}
