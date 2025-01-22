package views

import (
	"beverages-cli/handler"
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

		// # Exit condition
		if customerId == "0" {
			break
		}
	}
}
