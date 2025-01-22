package submenu1

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func SubMenuAdmin(admin entity.Users) {
	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Welcome", admin.Name)
		fmt.Println("=======================================")
		fmt.Println("")

		// # Fetch data customers
		customers, err := handler.GetAllCustomers()
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "USERNAME", "NAME", "EMAIL", "PHONE", "ADDRESS"}) // Table headers

		// # Add rows to the table
		for _, value := range customers {
			row := []string{
				strconv.Itoa(value.UserID),
				value.Username,
				value.Name,
				value.Email,
				value.PhoneNumber,
				value.Address,
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()

		// # UI Below Table
		fmt.Println("")
		fmt.Println("1. Register Customer")
		fmt.Println("2. Delete Customer")
		fmt.Println("3. Edit Customer")
		fmt.Println("4. Exit")
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
			AddCustomer()
		case "2":
			DeleteCustomer()
		case "3":
			EditCustomer()
		case "4":
		default:
			continue
		}

		// # Exit condition
		if userInput == "4" {
			break
		}
	}
}
