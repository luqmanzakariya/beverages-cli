package handler

import (
	"beverages-cli/config"
	"beverages-cli/entity"
	"context"
	"log"
	"time"
)

func GetUsersById(username string, password string) ([]entity.Users, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

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

func GetAllCustomers() ([]entity.Users, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

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

func CreateCustomer(Username string, Name string, Email string, PhoneNumber string, Address string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		INSERT INTO Users (Username, Password, Name, Email, PhoneNumber, Address, RoleID) VALUES 
		(?, NULL, ?, ?, ?, ?, 2);
	`

	rows, err := db.QueryContext(ctx, query, Username, Name, Email, PhoneNumber, Address)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func DeleteCustomerById(UserID string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		DELETE FROM Users 
		WHERE UserID = ?;
	`

	rows, err := db.QueryContext(ctx, query, UserID)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

func EditCustomer(UserId string, Username string, Name string, Email string, PhoneNumber string, Address string) (bool, error) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// # Query with parameters
	query := `
		UPDATE Users
		SET Username  = ?, Name = ?, Email = ?, PhoneNumber = ?, Address = ?
		WHERE UserID  = ?;
	`

	rows, err := db.QueryContext(ctx, query, Username, Name, Email, PhoneNumber, Address, UserId)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}
