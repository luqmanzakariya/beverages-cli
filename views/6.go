package views

import "fmt"

func View6() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("=======================================")
	fmt.Println("Good bye~")
	fmt.Println("=======================================")
}
