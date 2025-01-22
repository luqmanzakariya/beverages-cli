package subMenu1

import (
	config "beverages-cli/config"
	"beverages-cli/entity"
	"beverages-cli/handler"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func SubMenuAdmin(admin entity.Users) {
	// Init and close db after function ends
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	for {
		// # User interface
		fmt.Print("\033[H\033[2J")
		fmt.Println("=======================================")
		fmt.Println("Welcome", admin.Name)
		fmt.Println("=======================================")
		fmt.Println("")

		// # Fetch data customers
		customers, err := handler.GetAllCustomers(db)
		if err != nil {
			log.Fatal("Failed to retrieve data from database: ", err)
		}

		// # Create a new table writer
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "NAME", "PHONE"}) // Table headers

		// # Add rows to the table
		for _, value := range customers {
			row := []string{
				strconv.Itoa(value.UserID),
				value.Name,
				value.PhoneNumber,
			}
			table.Append(row)
		}

		// # Render the response and table
		table.Render()

		// # UI Above Table
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

		fmt.Println(userInput)

		// # Exit condition
		if userInput == "4" {
			break
		}
	}
}
