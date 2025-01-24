package entity

import (
	"database/sql"
	"time"
)

type Orders struct {
	OrderID     int
	UserID      string
	OrderDate   time.Time
	TotalAmount float32
	Status      string
}

type Cart struct {
	ID          sql.NullInt64
	ProductName sql.NullString
	Quantity    sql.NullInt64
	Price       sql.NullFloat64
	Total       sql.NullFloat64
}
