package views

import (
	// submenu4 "beverages-cli/views/submenu4"
	"fmt"
	"log"
)

func View4() {
	// submenu4.addCategories()
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
