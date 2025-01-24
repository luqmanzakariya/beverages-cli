package views

import (
	"beverages-cli/views/submenu5"
	"fmt"
	"log"
)

func View5() {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Reports")
		fmt.Println("=======================================")
		fmt.Println("")
		fmt.Println("1. List Registered Customer")
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

		// # Main View
		switch userInput {
		case "1":
			submenu5.ListRegisteredUser()
		case "2":
			submenu5.ListProductAndStock()
		case "3":
		case "4":
		case "5":
		case "6":
		default:
			continue
		}

		// # Exit condition
		if userInput == "6" {
			break
		}
	}
}
