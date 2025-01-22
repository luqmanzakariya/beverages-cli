package views

import (
	"fmt"
	"log"
)

func View2() {
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
