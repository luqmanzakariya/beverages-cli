package submenu1

import (
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func AddNewCategory(CategoryID int, CategoryName string) {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Add New Category")
		fmt.Println("=======================================")

		// # Fetch data list
		carts, err := handler.GetListCategories()
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "CATEGORY"}) // Table headers

		// # Add rows to the table
		for _, value := range carts {
			row := []string{
				strconv.Itoa(value.CategoryID),
				value.CategoryName,
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()

		// # Waiting product id
		fmt.Println("")
		fmt.Print("Please enter Category ID (input 0 to exit): ")
		var categoryID string
		_, err = fmt.Scan(&categoryID)
		if err != nil {
			log.Fatal("Failed to read productID")
		}

		// # Exit condition
		if categoryID == "0" {
			break
		}

		// # Waiting input quantity
		fmt.Println("")
		fmt.Print("Please enter product ID (input 0 to exit): ")
		var quantity string
		_, err = fmt.Scan(&quantity)
		if err != nil {
			log.Fatal("Failed to read productID")
		}

		// # Exit condition
		if quantity == "0" {
			break
		}
	}
}
