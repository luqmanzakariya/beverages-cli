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

func TopSalesPerCategory() ([]entity.TopSalesPerCategory, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT 
				ps.CategoryName,
				ps.ProductName,
				ps.TotalSales
		FROM 
				(
						SELECT 
								p.ProductID,
								p.ProductName,
								c.CategoryID,
								c.CategoryName,
								SUM(od.Quantity * od.Price) AS TotalSales
						FROM 
								OrderDetails od
						JOIN 
								Products p ON od.ProductID = p.ProductID
						JOIN 
								Categories c ON p.CategoryID = c.CategoryID
						GROUP BY 
								p.ProductID, p.ProductName, c.CategoryID, c.CategoryName
				) AS ps
		JOIN 
				(
						SELECT 
								CategoryID,
								MAX(TotalSales) AS MaxSales
						FROM 
								(
										SELECT 
												p.ProductID,
												c.CategoryID,
												SUM(od.Quantity * od.Price) AS TotalSales
										FROM 
												OrderDetails od
										JOIN 
												Products p ON od.ProductID = p.ProductID
										JOIN 
												Categories c ON p.CategoryID = c.CategoryID
										GROUP BY 
												p.ProductID, c.CategoryID
								) AS SubSales
						GROUP BY 
								CategoryID
				) AS TopSalesPerCategory
		ON 
				ps.CategoryID = TopSalesPerCategory.CategoryID
				AND ps.TotalSales = TopSalesPerCategory.MaxSales
		ORDER BY 
				ps.CategoryName;
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.TopSalesPerCategory
	for rows.Next() {
		var t entity.TopSalesPerCategory
		if err = rows.Scan(&t.CategoryName, &t.ProductName, &t.TopSales); err != nil {
			return nil, err
		}
		orders = append(orders, t)
	}

	return orders, nil
}

func TopSpenderCustomer() ([]entity.TopSpenderCustomer, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT 
				u.Name AS CustomerName,
				u.Email AS CustomerEmail,
				u.PhoneNumber AS CustomerPhone,
				SUM(o.TotalAmount) AS TotalSpent
		FROM 
				Orders o
		JOIN 
				Users u ON o.UserID = u.UserID
		WHERE 
				u.RoleID = (SELECT RoleID FROM Roles WHERE RoleName = 'Customer')
		GROUP BY 
				u.UserID, u.Name, u.Email
		ORDER BY 
				TotalSpent DESC;
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.TopSpenderCustomer
	for rows.Next() {
		var t entity.TopSpenderCustomer
		if err = rows.Scan(&t.CustomerName, &t.CustomerEmail, &t.CustomerPhone, &t.TotalSpent); err != nil {
			return nil, err
		}
		orders = append(orders, t)
	}

	return orders, nil
}
