package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func GetLisCategory() ([]entity.Categories, error) {
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
			CategoryName
		FROM Categories c
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

func CreateCategory(CategoryName string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		INSERT INTO categories (CategoryName) VALUES 
		(?, NULL, ?, ?, ?, ?, 2);
	`

	rows, err := db.QueryContext(ctx, query, CategoryName)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func DeleteCategoryById(CategoryID string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		DELETE FROM Categories 
		WHERE CategoryID = ?;
	`

	rows, err := db.QueryContext(ctx, query, CategoryID)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func EditCategory(CategoryID string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		UPDATE Categories
		SET CategoryName  = ?,
		WHERE CategoriesID  = ?;
	`

	rows, err := db.QueryContext(ctx, query, CategoryID)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}
