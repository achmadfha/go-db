package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func IsValidPhoneNumber(phoneNumber string) bool {

	phoneRegex := regexp.MustCompile(`^\d{1,12}$`)
	return phoneRegex.MatchString(phoneNumber)
}

func ValidationJumlah() string {
	for {

		fmt.Print("Enter Satuan (kg/buah) :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		jumlah := scanner.Text()

		if jumlah == "kg" || jumlah == "buah" {
			return jumlah
		}

		fmt.Println("Invalid input. Please enter 'kg' or 'buah'.")
	}
}

func IsValidNumber(price string) bool {

	numberRegex := regexp.MustCompile("^[0-9]+$")
	return numberRegex.MatchString(price)
}

func IsValidSatuan(satuan string) bool {
	return satuan == "kg" || satuan == "buah"
}

func ReadInput(prompt string, validator func(string) (interface{}, error)) (interface{}, error) {
	for {
		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		userInput := scanner.Text()

		result, err := validator(userInput)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		return result, nil
	}
}

func ValidateDate(input string) (interface{}, error) {

	minDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Now()

	parsedDate, err := time.Parse("02-01-2006", input)
	if err != nil || parsedDate.Before(minDate) || parsedDate.After(maxDate) {
		return nil, fmt.Errorf("please enter a valid date in dd-mm-yyyy format")
	}

	return parsedDate, nil
}
