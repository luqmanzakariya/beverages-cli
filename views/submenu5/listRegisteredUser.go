package submenu5

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func ListRegisteredUser() {
	// # User interface
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("List Registered Customer")
	fmt.Println("=======================================")
	fmt.Println("")

	// # Fetch all customers
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
	fmt.Println("")

	// # Waiting user input
	var userInput string
	_, err = fmt.Scan(&userInput)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
}
