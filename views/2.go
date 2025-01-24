package views

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"beverages-cli/views/submenu2"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func View2() {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Order")
		fmt.Println("=======================================")
		fmt.Println("")

		// # Fetch data customers
		customers, err := handler.GetAllCustomers()
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "USERNAME", "NAME", "EMAIL", "PHONE", "ADDRESS"}) // Table headers

		// # Add rows to the table
		for _, value := range customers {
			row := []string{
				strconv.Itoa(value.UserID),
				value.Username,
				value.Name,
				value.Email,
				value.PhoneNumber,
				value.Address,
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()

		// # Waiting input Customer ID
		fmt.Println("")
		fmt.Print("Choose registered customer ID to order (0 to exit): ")
		var customerId string
		_, err = fmt.Scan(&customerId)
		if err != nil {
			log.Fatal("Failed to read customerId")
		}

		var foundIdCustomers bool = false
		var selectedCustomers entity.Users
		for _, value := range customers {
			if strconv.Itoa(value.UserID) == customerId {
				foundIdCustomers = true
				selectedCustomers = value
			}
		}

		// # Exit condition
		if customerId == "0" {
			break
		} else {
			// # Go to sub menu if customers found
			if foundIdCustomers {
				orderId, orderDate, err := handler.CreateOrder(customerId)
				if err != nil {
					log.Fatal("Failed to retrieve data from database: ", err)
				}

				// # Sub Menu Order Product
				submenu2.CartOrder(orderId, orderDate, selectedCustomers)

			} else {
				// # Customers not found
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
}
