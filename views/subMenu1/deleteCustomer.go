package submenu1

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"fmt"
	"log"
	"strconv"
)

func DeleteCustomer(listCustomer []entity.Users) {
	for {
		fmt.Println("")
		fmt.Print("Select id customer you want to delete: ")
		var deletedId string
		_, err := fmt.Scan(&deletedId)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}
		var foundIdCustomers bool = false
		for _, value := range listCustomer {
			if strconv.Itoa(value.UserID) == deletedId {
				foundIdCustomers = true
			}
		}

		if foundIdCustomers {
			successDelete, err := handler.DeleteCustomerById(deletedId)
			if err != nil {
				log.Fatal("Failed delete customer by ID")
			}
			if successDelete {
				fmt.Println("")
				fmt.Println("Customer successfully deleted!")

				var exitInput string
				_, err = fmt.Scan(&exitInput)
				if err != nil {
					log.Fatal("Failed to read customerId")
				}
				break
			} else {
				fmt.Println("Error delete customer")
			}

			_, err = fmt.Scan(&deletedId)
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
