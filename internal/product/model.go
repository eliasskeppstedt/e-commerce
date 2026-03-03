package product

type Product struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Stock        int     `json:"stock"`
	Price        float64 `json:"price"`
	Manufacturer string  `json:"manufacturer"`
	Description  string  `json:"description"`
	CategoryID   int     `json:"category_id"`
}
