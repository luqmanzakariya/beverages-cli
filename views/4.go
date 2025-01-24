package views

import (
	// submenu4 "beverages-cli/views/submenu4"
	"beverages-cli/handler"
	// "beverages-cli/views/submenu4"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func View4() {
	// submenu4.addCategories()
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Categories")
		fmt.Println("=======================================")

		// # Fetch data customers
		categories, err := handler.GetLisCategory()
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "NAME"}) // Table headers

		// # Add rows to the table
		for _, value := range categories {
			row := []string{
				strconv.Itoa(value.CategoryID),
				value.CategoryName,
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()
		fmt.Println("")
		fmt.Println("1. Add new category")
		fmt.Println("2. Edit category")
		fmt.Println("3. Delete category")
		fmt.Println("4. Exit")
		fmt.Print("Answer: ")

		// # Waiting user input
		var userInput string
		_, err = fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		if userInput == "3" {
			// submenu4.DeleteCategory(categories)
			fmt.Println("")
			fmt.Print("Please select ID you want to delete: ")
			var selectedId string
			_, err = fmt.Scan(&selectedId)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			_, err := handler.DeleteCategoryById(selectedId)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}
		}

		// # Exit condition
		if userInput == "4" {
			break
		}
	}
}
