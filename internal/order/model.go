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
	OrderID int             `json:"order_id"`
	Status  string          `json:"status"`
	Date    string          `json:"date"`
	Items   []OrderItemView `json:"items"`
}

type OrderItemView struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type UserOrders struct {
	UserID   int              `json:"user_id"`
	Username string           `json:"username"`
	Orders   []OrderWithItems `json:"orders"`
}
