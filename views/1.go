package views

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	subMenu1 "beverages-cli/views/submenu1"
	"fmt"
	"log"
)

func View1() {
	// # Flag to check login or not
	var loggedIn bool = false
	var adminDetail entity.Users

	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Admin")
		fmt.Println("=======================================")
		fmt.Println("Please login as admin first")
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

		// # Fetching data from database
		users, err := handler.GetUsersById(username, password)
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Check if password is correct and RoleID is 1 (admin)
		if len(users) == 1 && users[0].Password.Valid {
			if users[0].Password.String == password && users[0].RoleID == 1 {
				loggedIn = true
				adminDetail = users[0]
			} else {
				fmt.Println("")
				fmt.Println("Wrong password!")
			}
		} else {
			fmt.Println("")
			fmt.Println("User not found")
		}

		// # Break the loop if successfully login
		if loggedIn {
			break
		} else {
			fmt.Print("Continue? (press 'n' to exit) ")
		}

		// # Input to retry or exit
		var actionInput string
		_, err = fmt.Scan(&actionInput)
		if err != nil {
			log.Fatal("Failed to read password")
		}

		// # Exit condition
		if actionInput == "n" {
			break
		}
	}

	// # Show admin menu if successfully logged in
	if loggedIn {
		// # Sub Menu Admin
		subMenu1.SubMenuAdmin(adminDetail)
	}
}
