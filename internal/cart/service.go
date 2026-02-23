package cart

import (
	"ecommerce/duckyarmy/internal/product"
	_ "errors"
	"fmt"
)

type cartService interface {
	cartItemAdd(carttID, productID int) error
	cartItemBulkAdd(cartID, productID int) error
	cartItemRemove(cartID, productID int) error
	cartItemBulkRemove(cartID, productID int) error
}

type CartService1 struct {
	productRepo product.ProductRepository
	cartRepo    CartRepository
}

func NewCartService1(pR product.ProductRepository, cR CartRepository) *CartService1 {
	return &CartService1{productRepo: pR, cartRepo: cR}
}

func (s *CartService1) cartItemAdd(cartID, productID int) error {
	stock, err := s.productRepo.GetProductStock(productID)

	if err != nil {
		return err
	}
	if stock == 0 {
		return fmt.Errorf("Product %d not in stock", productID)
	}

	return s.cartRepo.cartItemAdd(cartID, productID)
}

func (s *CartService1) cartItemBulkAdd(cartID, productID int) error {
	return nil
}

func (s *CartService1) cartItemRemove(cartID, productID int) error {
	return nil
}

func (s *CartService1) cartItemBulkRemove(cartID, productID int) error {
	return nil
}
