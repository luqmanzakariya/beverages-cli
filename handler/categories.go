package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func GetListCategories() ([]entity.Categories, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT
			CategoryID,
			CategoryName,
		FROM Categories c
		RIGHT JOIN Products p on c.ProductID = p.ProductID 
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Categories
	for rows.Next() {
		var c entity.Categories
		if err = rows.Scan(&c.CategoryID, &c.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
