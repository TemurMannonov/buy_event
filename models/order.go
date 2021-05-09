package models

type Order struct {
	ID         string  `json:"id" db:"id"`
	CustomerID string  `json:"customer_id" db:"customer_id"`
	Products   string  `json:"product" db:"products"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}
