package views

import (
	config "beverages-cli/config"
	"fmt"
	"log"
)

func View5() {
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
		fmt.Println("Reports")
		fmt.Println("=======================================")
		fmt.Println("")
		fmt.Println("1. List Registered User")
		fmt.Println("2. List Beverages and Stocks")
		fmt.Println("3. List Order")
		fmt.Println("4. Top Sales Per Category")
		fmt.Println("5. Top Spender Customer")
		fmt.Println("6. Exit")
		fmt.Print("Answer: ")

		// # Waiting user input
		var userInput string
		_, err := fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		// # Exit condition
		if userInput == "6" {
			break
		}
	}
}
