package views

import (
	config "beverages-cli/config"
	"fmt"
	"log"
)

func View1() {
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
		fmt.Println("Admin")
		fmt.Println("=======================================")
		fmt.Println("*input 0 to username and 0 to exit")
		fmt.Println("")

		// # Waiting input username
		fmt.Print("Please input username: ")
		var username string
		_, err := fmt.Scan(&username)
		if err != nil {
			log.Fatal("Failed to read username")
		}

		// # Waiting input password
		fmt.Print("Please input password: ")
		var password string
		_, err = fmt.Scan(&password)
		if err != nil {
			log.Fatal("Failed to read password")
		}

		// # Exit condition
		if username == "0" && password == "0" {
			break
		}
	}
}
