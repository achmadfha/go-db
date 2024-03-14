package handlers

import (
	"bufio"
	"challenge-godb/models"
	"challenge-godb/services"
	"challenge-godb/utils"
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"os"
	"strconv"
	"strings"
)

func ViewAllServices() {
	pelayanans, err := services.GetAllServices()
	if err != nil {
		log.Fatal(err)
	}

	if len(pelayanans) == 0 {
		fmt.Println("No Service Found Sorry :(")
		return
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "No"},
			{Text: "Pelayanan"},
			{Text: "Jumlah"},
			{Text: "Harga"},
		},
	}

	for i, pelayanan := range pelayanans {
		row := []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i+1)},
			{Text: pelayanan.JenisPelayanan},
			{Text: pelayanan.Satuan},
			{Text: fmt.Sprintf("%d", pelayanan.Harga)},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	fmt.Println("\n All Services")
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func AddServices() {

	fmt.Print("\nEnter Pelayanan name: ")
	scanner := bufio.NewScanner(os.Stdin)
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

	newService := models.Pelayanan{
		JenisPelayanan: servicesName,
		Satuan:         pelayananJumlah,
		Harga:          prices,
	}

	fmt.Print("Are you sure you want to add this customer? (y/n): ")

	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		services.AddNewServices(newService)
		fmt.Println("New Services added successfully!\n")
	} else {
		fmt.Println("Customer addition canceled.\n")
	}
}

func UpdateSevices() {
	fmt.Print("\nEnter Services ID: ")
	var pelayananID int
	fmt.Scanln(&pelayananID)

	existingServices, err := services.ViewServicesByID(pelayananID)
	if err != nil {
		log.Println("Error checking customer existence:", err)
		return
	}

	if len(existingServices) == 0 {
		fmt.Println("Service with ID", pelayananID, "not found.")
		return
	}

	updatedService := readServiceDetails(existingServices[0])

	fmt.Print("Are you sure you want to update this service? (y/n): ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) == "y" {
		// Update the service
		err := services.UpdateServicesByID(pelayananID, updatedService)
		if err != nil {
			log.Println("Error updating service:", err)
			return
		}
		fmt.Println("Service updated successfully!")
	} else {
		fmt.Println("Service update canceled.")
	}
}

func readServiceDetails(existingService models.Pelayanan) models.Pelayanan {
	var updatedService models.Pelayanan

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Updated Service Name (%s): ", existingService.JenisPelayanan)
	serviceName, _ := reader.ReadString('\n')
	serviceName = strings.TrimSpace(serviceName)
	if serviceName == "" {
		updatedService.JenisPelayanan = existingService.JenisPelayanan
	} else {
		updatedService.JenisPelayanan = serviceName
	}

	for {
		fmt.Printf("Enter Updated Jumlah (%s):", existingService.Satuan)
		satuanJumlah, _ := reader.ReadString('\n')
		satuanJumlah = strings.TrimSpace(satuanJumlah)

		if satuanJumlah == "" {
			updatedService.Satuan = existingService.Satuan
			break
		} else {
			if utils.IsValidSatuan(satuanJumlah) {
				updatedService.Satuan = satuanJumlah
				break
			} else {
				fmt.Println("Invalid Satuan! Please enter 'kg' or 'buah'.")
			}
		}
	}

	// Validate and read updated price
	var price string
	fmt.Printf("Enter Updated Price (%d): ", existingService.Harga)
	fmt.Scanln(&price)
	if price == "" {
		updatedService.Harga = existingService.Harga
	} else {
		for {
			if utils.IsValidNumber(price) {
				updatedService.Harga, _ = strconv.Atoi(price)
				break
			}
			fmt.Println("Invalid Price! Please enter a valid number.")
		}
	}

	return updatedService
}
