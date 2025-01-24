package submenu4

import (
	"beverages-cli/entity"
	"beverages-cli/handler"
	"fmt"
	"log"
	"strconv"
)

func DeleteCategory(listCategory []entity.Categories) {
	for {
		fmt.Println("")
		fmt.Print("Please select category you want to delete: ")
		var deletedId string
		_, err := fmt.Scan(&deletedId)
		if err != nil {
			log.Fatal("Failed to read userInput")
		}
		var foundIdCategories bool = false
		for _, value := range listCategory {
			if strconv.Itoa(value.CategoryID) == deletedId {
				foundIdCategories = true
			}
		}

		if foundIdCategories {
			successDelete, err := handler.DeleteCategoryById(deletedId)
			if err != nil {
				log.Fatal("Failed delete category by ID")
			}
			if successDelete {
				fmt.Println("")
				fmt.Println("Cateogory successfully deleted!")

				var exitInput string
				_, err = fmt.Scan(&exitInput)
				if err != nil {
					log.Fatal("Failed to read categoryId")
				}
				break
			} else {
				fmt.Println("Error delete category")
			}

			_, err = fmt.Scan(&deletedId)
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
