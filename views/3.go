package views

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func View3() {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("List Product & Stock")
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
		fmt.Println("1. Add new product")
		fmt.Println("2. Edit product")
		fmt.Println("3. Delete product")
		fmt.Println("4. Exit")
		fmt.Print("Answer: ")

		// # Waiting user input
		var userInput string
		_, err = fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		// # Exit condition
		if userInput == "4" {
			break
		}
	}
}
