package submenu1

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func EditCustomer(listCustomer []entity.Users) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// # Store Input
		fmt.Println("")
		fmt.Print("Select id customer you want to edit: ")
		var UserId string
		_, err := fmt.Scan(&UserId)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		var foundIdCustomers bool = false
		for _, value := range listCustomer {
			if strconv.Itoa(value.UserID) == UserId {
				foundIdCustomers = true
			}
		}

		if foundIdCustomers {
			// # UI and Store Input
			fmt.Print("\033[H\033[2J")
			fmt.Println("=======================================")
			fmt.Println("Edit Customer")
			fmt.Println("=======================================")
			fmt.Println("")

			fmt.Print("Input Username: ")
			Username, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			fmt.Print("Input Name: ")
			Name, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			fmt.Print("Input Email: ")
			Email, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			fmt.Print("Input Phone: ")
			PhoneNumber, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			fmt.Print("Input Address: ")
			Address, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			// # Create Customer
			customers, err := handler.EditCustomer(UserId, Username, Name, Email, PhoneNumber, Address)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}
			if customers {
				fmt.Println("")
				fmt.Println("Customer successfully editted!")

				var exitInput string
				_, err = fmt.Scan(&exitInput)
				if err != nil {
					log.Fatal("Failed to read customerId")
				}
				break
			} else {
				fmt.Println("Error create customer")
			}

			var wait string
			_, err = fmt.Scan(&wait)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}
		} else {
			fmt.Println("")
			fmt.Println("Customer not found (0 to exit)")
			var exitInput string
			_, err = fmt.Scan(&exitInput)
			if err != nil {
				log.Fatal("Failed to read customerId")
			}

			if exitInput == "0" {
				break
			}
		}
	}

}
