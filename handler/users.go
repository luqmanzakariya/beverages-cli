package handler

import (
	"beverages-cli/entity"
	"context"
	"database/sql"
	"time"
)

func GetUsersById(db *sql.DB, username string, password string) ([]entity.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		SELECT *
		FROM Users u 
		WHERE Username = ?
	`

	rows, err := db.QueryContext(ctx, query, username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		var u entity.Users

		// # Scan the result into the struct fields
		if err = rows.Scan(&u.UserID, &u.Username, &u.Password, &u.Name, &u.Email, &u.PhoneNumber, &u.Address, &u.RoleID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetAllCustomers(db *sql.DB) ([]entity.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		SELECT *
		FROM Users u
		WHERE u.RoleID = 2
	`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		var u entity.Users

		// # Scan the result into the struct fields
		if err = rows.Scan(&u.UserID, &u.Username, &u.Password, &u.Name, &u.Email, &u.PhoneNumber, &u.Address, &u.RoleID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
