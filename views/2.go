package views

import (
	config "beverages-cli/config"
	"fmt"
	"log"
)

func View2() {
	// Init and close db after function ends
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Order")
		fmt.Println("=======================================")
		fmt.Println("*input 0 to exit")
		fmt.Println("")

		// # Waiting input Customer ID
		fmt.Print("Please choose customer number: ")
		var customerId string
		_, err := fmt.Scan(&customerId)
		if err != nil {
			log.Fatal("Failed to read customerId")
		}

		// # Exit condition
		if customerId == "0" {
			break
		}
	}
}
