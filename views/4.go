package views

import (
	config "beverages-cli/config"
	"fmt"
	"log"
)

func View4() { 
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
		fmt.Println("Categories")
		fmt.Println("=======================================")
		fmt.Println("")
		fmt.Println("1. Add new category")
		fmt.Println("2. Edit category")
		fmt.Println("3. Delete category")
		fmt.Println("4. Exit")
		fmt.Print("Answer: ")

		// # Waiting user input
		var userInput string
		_, err := fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		// # Exit condition
		if userInput == "4" {
			break
		}
	}
}
