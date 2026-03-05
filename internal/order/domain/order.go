package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID         string
	UserID     string
	TotalPrice decimal.Decimal
	Status     string // e.g., "pending", "completed", "cancelled"
	CreatedAt  time.Time
	Items      []OrderItem
}

type OrderItem struct {
	ID          string
	OrderID     string
	ProductID   string
	ProductName string
	Price       decimal.Decimal
	Quantity    int
}

type OrderRepository interface {
	CreateOrder(order *Order) error
	UpdateProductStock(productID string, quantity int) error
}
