package submenu4

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func EditCategory(listCategory []entity.Users) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// # Store Input
		fmt.Println("")
		fmt.Print("Please select category you want to delete")
		var CategoryId string
		_, err := fmt.Scan(&CategoryId)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}

		var foundIdCategory bool = false
		for _, value := range listCategory {
			if strconv.Itoa(value.UserID) == CategoryId {
				foundIdCategory = true
			}
		}

		if foundIdCategory {
			// # UI and Store Input
			fmt.Print("\033[H\033[2J")
			fmt.Println("=======================================")
			fmt.Println("Edit Category")
			fmt.Println("=======================================")
			fmt.Println("")

			fmt.Print("Input CategoryID: ")
			CategoryID, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("Failed to read userInput")
			}

			// # edit category

			categories, err := handler.EditCategory(CategoryName)
	if err != nil {
		log.Fatal("Failed to read categoryID")
	}
	if categories {
		fmt.Println("")
		fmt.Println("Category successfully edited!")

				var exitInput string
				_, err = fmt.Scan(&exitInput)
				if err != nil {
					log.Fatal("Failed to read categoryId")
				}
				break
			} else {
				fmt.Println("Error create category")
			}

			var wait string
			_, err = fmt.Scan(&wait)
			if err != nil {
				log.Fatal("Failed to read userInput")
			}
		} else {
			fmt.Println("")
			fmt.Println("Category not found (0 to exit)")
			var exitInput string
			_, err = fmt.Scan(&exitInput)
			if err != nil {
				log.Fatal("Failed to read categoryId")
			}

			if exitInput == "0" {
				break
			}
		}
	}

}
