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

		// # Fetch data list cart
		carts, err := handler.GetListProducts()
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "PRODUCT NAME", "PRICE", "AVAILABLE", "CATEGORY"}) // Table headers

		// # Add rows to the table
		for _, value := range carts {
			row := []string{
				strconv.Itoa(value.ProductID),
				value.ProductName,
				fmt.Sprintf("%f", value.Price),
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

		// # Exit condition
		if productID == "0" {
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
