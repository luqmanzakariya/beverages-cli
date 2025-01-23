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

func CartOrder(OrderID int, orderDate string, customer entity.Users) {
	for {
		// # Initiate total variable
		var grandTotal float32 = 0
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Welcome", customer.Name)
		fmt.Println("Order ID:", OrderID)
		fmt.Println("Date:", orderDate)
		fmt.Println("==============List Cart================")

		// # Fetch data list cart
		carts, err := handler.GetListCart(OrderID)
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "PRODUCT NAME", "QUANTITY", "PRICE", "TOTAL"}) // Table headers

		// # Add rows to the table
		for _, value := range carts {
			grandTotal += float32(value.Total.Float64)
			row := []string{
				strconv.Itoa(int(value.ID.Int64)),
				value.ProductName.String,
				strconv.Itoa(int(value.Quantity.Int64)),
				fmt.Sprintf("%.2f", value.Price.Float64),
				fmt.Sprintf("%.2f", value.Total.Float64),
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()
		fmt.Println("TOTAL: ", grandTotal)

		// # UI Below Table
		fmt.Println("")
		fmt.Println("1. Add items")
		fmt.Println("2. Delete items")
		fmt.Println("3. Pay")
		fmt.Println("4. Cancel")
		fmt.Print("Answer: ")

		// # Waiting user input
		var userInput string
		_, err = fmt.Scan(&userInput)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		// # Submenu View
		switch userInput {
		case "1":
			AddItemsToCart(OrderID, orderDate, customer)
		case "2":

		case "3":

		case "4":
		default:
			continue
		}

		// # Exit condition
		if userInput == "3" || userInput == "4" {
			break
		}
	}
}
