package product

import (
	"ecommerce/duckyarmy/internal/review"
)

type Product struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Stock        int     `json:"stock"`
	Price        float64 `json:"price"`
	Manufacturer string  `json:"manufacturer"`
	Description  string  `json:"description"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
}

type ProductCard struct {
	Product Product         `json:"product_id"`
	Review  []review.Review `json:"reviews"`
}
