package submenu1

import (
	"beverages-cli/handler"
	"fmt"
	"log"
)

func DeleteCustomer() {
	fmt.Println("")
	fmt.Print("Select id customer you want to delete: ")
	var deletedId string
	_, err := fmt.Scan(&deletedId)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
	successDelete, err := handler.DeleteCustomerById(deletedId)
	if err != nil {
		log.Fatal("Failed delete customer by ID")
	}
	if successDelete {
		fmt.Println("Customer successfully deleted")
	} else {
		fmt.Println("Error delete customer")
	}

	_, err = fmt.Scan(&deletedId)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
}
