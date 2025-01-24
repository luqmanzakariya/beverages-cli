package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func GetListOrder() ([]entity.OrderReports, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT 
			OrderID,
			u.Name AS Name,
			OrderDate,
			TotalAmount,
			Status
		FROM Orders o 
		LEFT JOIN Users u ON o.UserID = u.UserID 
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrderReports
	for rows.Next() {
		var c entity.OrderReports
		if err = rows.Scan(&c.OrderID, &c.Name, &c.OrderDate, &c.TotalAmount, &c.Status); err != nil {
			return nil, err
		}
		orders = append(orders, c)
	}

	return orders, nil
}
