package submenu5

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func ListProductAndStock() {
	// # User interface
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("List Product and Stock")
	fmt.Println("=======================================")
	fmt.Println("")

	// # Fetch data products
	products, err := handler.GetListProducts()
	if err != nil {
		log.Fatal("Failed to retrieve data from database: ", err)
	}

	// # Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "PRODUCT NAME", "PRICE", "AVAILABLE", "CATEGORY"}) // Table headers

	// # Add rows to the table
	for _, value := range products {
		row := []string{
			strconv.Itoa(value.ProductID),
			value.ProductName,
			fmt.Sprintf("%.2f", value.Price),
			strconv.Itoa(value.StockQuantity),
			value.CategoryName,
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
