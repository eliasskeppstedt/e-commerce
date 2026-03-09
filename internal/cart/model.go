package cart

type Cart struct {
	CartID int    `json:"cart_id"`
	UserID int    `json:"user_id"`
	Status string `json:"status"`
}

type CartItem struct {
	CartItemID int `json:"cart_item_id"`
	CartID     int `json:"cart_id"`
	ProductID  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type AddToCartRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartItemRequest struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"product_price"`
	Subtotal    float64 `json:"subtotal"`
}
