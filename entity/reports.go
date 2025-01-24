package entity

type OrderReports struct {
	OrderID     int
	Name        string
	OrderDate   string
	TotalAmount float64
	Status      string
}

type TopSalesPerCategory struct {
	CategoryName string
	ProductName  string
	TopSales     float64
}

type TopSpenderCustomer struct {
	CustomerName  string
	CustomerEmail string
	CustomerPhone string
	TotalSpent    float64
}
