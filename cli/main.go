package main

import (
	"beverages-cli/views"
	"fmt"
	"log"
)

type Developers struct {
	DeveloperID int
	StudioName  string
	Location    string
}

func main() {
	for {
		// # Show Initial View
		views.InitialMenu()

		// # Waiting Initial Input
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal("Failed to read name input")
		}

		// # Main View
		switch input {
		case "1":
			views.View1()
		case "2":
			views.View2()
		case "3":
			views.View3()
		case "4":
			views.View4()
		case "5":
			views.View5()
		case "6":
			views.View6()
		default:
			continue
		}

		// # Exit menu if 6 is selected
		if input == "6" {
			break
		}
	}
}
