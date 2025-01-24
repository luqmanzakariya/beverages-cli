package submenu5

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func ListOrders() {
	// # User interface
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("List Order By Customer")
	fmt.Println("=======================================")
	fmt.Println("")

	// # Fetch all customers
	customers, err := handler.GetListOrder()
	if err != nil {
		log.Fatal("Failed to retrieve data from database: ", err)
	}

	// # Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Date", "Total", "Status"}) // Table headers

	// # Add rows to the table
	for _, value := range customers {
		row := []string{
			strconv.Itoa(value.OrderID),
			value.Name,
			value.OrderDate,
			strconv.FormatFloat(value.TotalAmount, 'f', 0, 64),
			// fmt.Sprintf("%.2f", value.TotalAmount),
			value.Status,
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
