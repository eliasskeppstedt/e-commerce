package order

import (
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/product"
	"errors"
	"fmt"
)

type OrderService interface {
	CheckOut(userID int) error
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
	// här börjar en transaktion, bör följa ACID.......
	cart, err := s.cartRepo.GetCartByUserID(userID)
	fmt.Println(cart.UserID)

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

	orderItems := make([]OrderItem, len(cartItems))

	for i, cartItem := range cartItems {
		fmt.Println(cartItem.Quantity, cartItem.ProductID)
		orderItems[i].Quantity = cartItem.Quantity
		orderItems[i].ProductID = cartItem.ProductID
		price := 1.5 //err, price := s.productRepo.GetPrice(cartItem.ProductID)
		fmt.Println("orderService1 CheckOut: hårdkodat pris, väntar på implementering i product")
		orderItems[i].PriceAtPurchase = price
	}
	fmt.Println(cart.UserID, orderItems[0].Quantity, orderItems[0].PriceAtPurchase, orderItems[0].ProductID)
	err = s.orderRepo.CheckOut(cart.UserID, orderItems)

	return err
}
