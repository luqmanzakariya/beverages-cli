package views

import "fmt"

func InitialMenu() {
	// Initial Interface
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("Welcome to Team 2 Wine & Beverages")
	fmt.Println("=======================================")
	fmt.Println("Please choose:")
	fmt.Println("1. Admin Menu")
	fmt.Println("2. Order")
	fmt.Println("3. Stock")
	fmt.Println("4. Categories")
	fmt.Println("5. Generate Reports")
	fmt.Println("6. Exit")
	fmt.Println("")
	fmt.Print("Answer: ")
}

func ReturnString(str string) string {
	return str
}
