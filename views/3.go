package views

import (
	"fmt"
	"log"
)

func View3() {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("List Product & Stock")
		fmt.Println("=======================================")
		fmt.Println("")
		fmt.Println("1. Add new product")
		fmt.Println("2. Edit product")
		fmt.Println("3. Delete product")
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
