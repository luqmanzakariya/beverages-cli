package entity

type OrderDetails struct {
	OrderDetailID int
	OrderID       int
	ProductID     int
	Quantity      int
	Price         float32
}