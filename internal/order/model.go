package order

import "time"

type Order struct {
	OrderID   int       `json:"order_id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	OrderItemID     int     `json:"order_item_id"`
	OrderID         int     `json:"order_id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

type OrderWithItems struct {
	OrderID int
	Date    string
	Status  string
	Items   []OrderItemWithProduct
}

type OrderItemWithProduct struct {
	ProductName string
	Quantity    int
	Price       float64
}
