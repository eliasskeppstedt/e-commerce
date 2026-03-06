package cart

import (
	"ecommerce/duckyarmy/internal/product"
	_ "errors"
	"fmt"
)

type CartService interface {
	AddItem(userID, productID, quantity int) error
	CreateCart(userID int) error
	GetCartProducts(userID int) ([]CartProduct, error)
	RemoveItem(userID, productID, quantity int) error
}

type cartService1 struct {
	productRepo product.ProductRepository
	cartRepo    CartRepository
}

func NewCartService1(pR product.ProductRepository, cR CartRepository) *cartService1 {
	return &cartService1{productRepo: pR, cartRepo: cR}
}

func (s *cartService1) AddItem(userID, productID, quantity int) error {

	stock, err := s.productRepo.GetProductStock(productID)
	if err != nil {
		return err
	}

	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	items, err := s.cartRepo.GetItems(cart.CartID)
	if err != nil {
		return err
	}

	curQuantity := 0

	for _, item := range items {
		if item.ProductID == productID {
			curQuantity = item.Quantity
		}
	}

	if curQuantity+quantity > stock {
		return fmt.Errorf("not enough stock")
	}
	err = s.cartRepo.AddItem(cart.CartID, productID, quantity)
	if err != nil {
		return err
	}

	err = s.productRepo.UpdateStock(productID, quantity)

	return err
}

func (s *cartService1) GetCartByUserID(userID int) (*Cart, error) {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *cartService1) CreateCart(userID int) error {
	return s.cartRepo.CreateCart(userID)
}

func (s *cartService1) GetCartProducts(userID int) ([]CartProduct, error) {

	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	items, err := s.cartRepo.GetItems(cart.CartID)
	if err != nil {
		return nil, err
	}

	var cartProducts []CartProduct

	for _, item := range items {

		p, err := s.productRepo.GetByProductID(item.ProductID)
		if err != nil {
			return nil, err
		}

		cartProducts = append(cartProducts, CartProduct{
			ProductID:   p.ProductID,
			ProductName: p.ProductName,
			Price:       p.Price,
			Quantity:    item.Quantity,
		})
	}

	return cartProducts, nil
}

func (s *cartService1) RemoveItem(userID, productID, quantity int) error {

	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	err = s.cartRepo.RemoveItem(cart.CartID, productID, quantity)
	if err != nil {
		return err
	}

	// returnera produkten till lagret
	err = s.productRepo.IncreaseStock(productID, quantity)
	if err != nil {
		return err
	}

	return nil
}
