package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func PaidOrder(OrderID int, grandTotal float32) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		UPDATE Orders
		SET Status = 'PAID', TotalAmount = ?
		WHERE OrderID = ?;
	`

	rows, err := db.QueryContext(ctx, query, grandTotal, OrderID)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func CreateOrder(UserID string) (int, string, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		INSERT INTO Orders (UserID, OrderDate) VALUES 
		(?, ?);
	`

	// # Get Current Date
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")

	// # Insert Process
	result, err := db.ExecContext(ctx, query, UserID, formattedTime)
	if err != nil {
		return 0, formattedTime, err
	}

	// Get Last Inserted ID
	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, formattedTime, err
	}

	return int(orderID), formattedTime, nil
}

func GetListCart(OrderID int) ([]entity.Cart, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		SELECT
			od.OrderDetailID  AS ID,
			p.ProductName AS ProductName,
			od.Quantity AS Quantity,
			p.Price AS Price,
			od.Price AS Total
		FROM Orders o
		LEFT JOIN OrderDetails od ON o.OrderID = od.OrderID 
		LEFT JOIN Products p ON od.ProductID = p.ProductID
		WHERE o.OrderID = ?
	`

	rows, err := db.QueryContext(ctx, query, OrderID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []entity.Cart
	for rows.Next() {
		var u entity.Cart

		// # Scan the result into the struct fields
		if err = rows.Scan(&u.ID, &u.ProductName, &u.Quantity, &u.Price, &u.Total); err != nil {
			return nil, err
		}

		if u.ID.Valid {
			carts = append(carts, u)
		}
	}

	return carts, nil
}
