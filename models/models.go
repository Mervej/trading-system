package models

import "time"

type User struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
}

type OrderStatus string

const (
	Accepted  OrderStatus = "ACCEPTED"
	Rejected  OrderStatus = "REJECTED"
	Completed OrderStatus = "COMPLETED"
)

type OrderType string

const (
	Buy  OrderType = "BUY"
	Sell OrderType = "SELL"
)

type Order struct {
	ID        string
	UserID    string
	Type      OrderType
	Symbol    string
	Quantity  int
	Price     float64
	Timestamp time.Time
	Status    OrderStatus
}

type Trade struct {
	ID            string
	Type          OrderType
	BuyerOrderID  string
	SellerOrderID string
	Symbol        string
	Quantity      int
	Price         float64
	Timestamp     time.Time
}
