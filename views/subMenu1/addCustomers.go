package submenu1

import (
	"beverages-cli/handler"
	"fmt"
	"log"
)

func AddCustomer() {
	// # UI and Store Input
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("Register New Customer")
	fmt.Println("=======================================")
	fmt.Println("")

	fmt.Print("Input Username: ")
	var Username string
	_, err := fmt.Scan(&Username)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	fmt.Print("Input Name: ")
	var Name string
	_, err = fmt.Scan(&Name)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	fmt.Print("Input Email: ")
	var Email string
	_, err = fmt.Scan(&Email)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	fmt.Print("Input Phone: ")
	var PhoneNumber string
	_, err = fmt.Scan(&PhoneNumber)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	fmt.Print("Input Address: ")
	var Address string
	_, err = fmt.Scan(&Address)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	// # Create Customer
	customers, err := handler.CreateCustomer(Username, Name, Email, PhoneNumber, Address)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
	if customers {
		fmt.Println("Customer successfully created")
	} else {
		fmt.Println("Error create customer")
	}

	var wait string
	_, err = fmt.Scan(&wait)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
}
