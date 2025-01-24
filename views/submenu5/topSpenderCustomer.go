package submenu5

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func TopSpenderCustomer() {
	// # User interface
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("Top Spender Customer")
	fmt.Println("=======================================")
	fmt.Println("")

	// # Fetch all customers
	customers, err := handler.TopSpenderCustomer()
	if err != nil {
		log.Fatal("Failed to retrieve data from database: ", err)
	}

	// # Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "EMAIL", "PHONE", "TOTAL SPENT"}) // Table headers

	// # Add rows to the table
	for _, value := range customers {
		row := []string{
			value.CustomerName,
			value.CustomerEmail,
			value.CustomerPhone,
			strconv.FormatFloat(value.TotalSpent, 'f', 0, 64),
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
