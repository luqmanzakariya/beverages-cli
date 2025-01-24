package submenu4

import (
	"beverages-cli/handler"
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateCategory() {
	reader := bufio.NewReader(os.Stdin)

	// # UI and Store Input
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("Please input category name")
	fmt.Println("=======================================")
	fmt.Println("")

	fmt.Print("Input Category Name: ")
	CategoryName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read userInput")
	}

	// # Create Category
	categories, err := handler.CreateCategory(CategoryName)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
	if categories {
		fmt.Println("")
		fmt.Println("Category successfully added!")
	} else {
		fmt.Println("Error create category")
	}

	var wait string
	_, err = fmt.Scan(&wait)
	if err != nil {
		log.Fatal("Failed to read userInput")
	}
}