package order

import (
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/product"
	"errors"
)

type OrderService interface {
	CheckOut(userID int) error
	GetOrders(userID int) ([]OrderWithItems, error)
}

type orderService1 struct {
	orderRepo   OrderRepository
	cartRepo    cart.CartRepository
	productRepo product.ProductRepository
}

func NewOrderService1(
	orderRepo OrderRepository,
	cartRepo cart.CartRepository,
	productRepo product.ProductRepository) *orderService1 {
	return &orderService1{orderRepo: orderRepo, cartRepo: cartRepo, productRepo: productRepo}
}

func (s *orderService1) CheckOut(userID int) error {

	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	cartItems, err := s.cartRepo.GetItems(cart.CartID)
	if err != nil {
		return err
	}

	if len(cartItems) == 0 {
		return errors.New("cannot checkout empty order")
	}

	// skapa order EN gång
	orderID, err := s.orderRepo.CreateOrder(userID)
	if err != nil {
		return err
	}

	for _, cartItem := range cartItems {

		err = s.orderRepo.AddOrderItem(OrderItem{
			OrderID:         orderID,
			ProductID:       cartItem.ProductID,
			Quantity:        cartItem.Quantity,
			PriceAtPurchase: 1.5,
		})

		if err != nil {
			return err
		}
	}

	err = s.cartRepo.SetCartInactive(cart.CartID)
	if err != nil {
		return err
	}

	return s.cartRepo.CreateCart(userID)
}

func (s *orderService1) GetOrders(userID int) ([]OrderWithItems, error) {
	return s.orderRepo.GetOrdersByUser(userID)
}
