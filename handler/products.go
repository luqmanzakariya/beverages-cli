package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func GetListProducts() ([]entity.Products, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT
			ProductID,
			ProductName,
			Price,
			StockQuantity,
			CategoryName
		FROM Products p
		LEFT JOIN Categories c ON p.CategoryID = c.CategoryID 
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Products
	for rows.Next() {
		var p entity.Products
		if err = rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.StockQuantity, &p.CategoryName); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
