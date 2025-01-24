package submenu2

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

func AddItemsToCart(OrderID int, orderDate string, customer entity.Users) {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Welcome", customer.Name)
		fmt.Println("Order ID:", OrderID)
		fmt.Println("Date:", orderDate)
		fmt.Println("=============List Products=============")

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

		// # Waiting product id
		fmt.Println("")
		fmt.Print("Please enter product ID (input 0 to exit): ")
		var productID string
		_, err = fmt.Scan(&productID)
		if err != nil {
			log.Fatal("Failed to read productID")
		}

		var foundProductID bool = false
		for _, value := range products {
			if strconv.Itoa(value.ProductID) == productID {
				foundProductID = true
			}
		}

		// # Exit condition
		if productID == "0" {
			break
		}

		// # Waiting input quantity
		fmt.Println("")
		fmt.Print("Please enter enter quantity (input 0 to exit): ")
		var quantity string
		_, err = fmt.Scan(&quantity)
		if err != nil {
			log.Fatal("Failed to read productID")
		}

		// # Exit condition
		if quantity == "0" {
			break
		}

		if !foundProductID {
			fmt.Println("")
			fmt.Print("Oops product not found")

			var continueInput string
			_, err = fmt.Scan(&continueInput)
			if err != nil {
				log.Fatal("Failed to read continueInput")
			}
		} else {
			successCreateProduct, err := handler.InsertProducts(OrderID, productID, quantity)
			if err != nil {
				log.Fatal("Failed to retrieve data from database: ", err)
			}

			if successCreateProduct {
				fmt.Println("")
				fmt.Println("Successfully added product")
			} else {
				fmt.Println("")
				fmt.Println("Error create product")
			}

			var continueInput string
			_, err = fmt.Scan(&continueInput)
			if err != nil {
				log.Fatal("Failed to read continueInput")
			}

			break
		}
	}
}
