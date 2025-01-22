package main

import (
	config "beverages-cli/config"
	"log"
)

type Developers struct {
	DeveloperID int
	StudioName  string
	Location    string
}

func main() {
	// Init and close db after function ends
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

}
