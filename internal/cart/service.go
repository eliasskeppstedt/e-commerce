package cart

import (
	"ecommerce/duckyarmy/internal/product"
	_ "errors"
	"fmt"
)

type CartService interface {
	AddItem(carttID, productID, quantity int) error
}

type cartService1 struct {
	productRepo product.ProductRepository
	cartRepo    CartRepository
}

func NewCartService1(pR product.ProductRepository, cR CartRepository) *cartService1 {
	return &cartService1{productRepo: pR, cartRepo: cR}
}

func (s *cartService1) AddItem(cartID, productID, quantity int) error {
	stock, err := s.productRepo.GetProductStock(productID)

	if err != nil {
		return err
	}
	if stock == 0 {
		return fmt.Errorf("Product %d not in stock", productID)
	}

	return s.cartRepo.AddItem(cartID, productID, quantity)
}
