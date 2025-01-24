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
			if len(carts) == 0 {
				fmt.Println("You don't have any order on the list")
			} else {
				fmt.Println("")
				fmt.Print("Please select order detail ID you want to delete? ")

				var orderDetailID string
				_, err = fmt.Scan(&orderDetailID)
				if err != nil {
					log.Fatal("Failed to read userInput")
				}

				var foundOrderDetailID bool = false
				for _, value := range carts {
					if strconv.Itoa(int(value.ID.Int64)) == orderDetailID {
						foundOrderDetailID = true
					}
				}

				if foundOrderDetailID {
					successDelete, err := handler.DeleteOrderDetailByID(orderDetailID)
					if err != nil {
						log.Fatal("Failed delete customer by ID")
					}

					if successDelete {
						fmt.Println("")
						fmt.Printf("Order detail with ID %s successfully deleted\n", orderDetailID)
					} else {
						fmt.Println("")
						fmt.Println("Error delete from database")
					}

				} else {
					fmt.Println("")
					fmt.Println("Oops, We dont find this order detail id in your cart")
				}
			}

			var userInputToContinue string
			_, err = fmt.Scan(&userInputToContinue)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

		case "3":
			successPaid, err := handler.PaidOrder(OrderID, grandTotal)
			if err != nil {
				log.Fatal("Failed delete customer by ID")
			}

			if successPaid {
				fmt.Println("")
				fmt.Println("Order with id: ", OrderID, " has been paid")
			} else {
				fmt.Println("Error paid order")
			}

			var userInputToContinue string
			_, err = fmt.Scan(&userInputToContinue)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

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
