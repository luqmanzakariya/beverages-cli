package entity

import (
	"database/sql"
)

type Users struct {
	UserID      int
	Username    string
	Password    sql.NullString // Allow string or null for the Password field
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	RoleID      int
}
