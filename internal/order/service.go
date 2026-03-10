package order

import (
	"context"
	"ecommerce/duckyarmy/internal/cart"
	"ecommerce/duckyarmy/internal/product"
	"ecommerce/duckyarmy/internal/transaction"
	"errors"
)

type OrderService interface {
	CheckOut(ctx context.Context, userID int) error
	GetOrdersByUser(ctx context.Context, userID int) ([]OrderWithItems, error)
	GetAllOrders(ctx context.Context) ([]UserOrders, error)
}

type orderService1 struct {
	tm          transaction.TxManager
	orderRepo   OrderRepository
	cartRepo    cart.CartRepository
	productRepo product.ProductRepository
}

func NewOrderService1(
	tm transaction.TxManager,
	orderRepo OrderRepository,
	cartRepo cart.CartRepository,
	productRepo product.ProductRepository,
) *orderService1 {
	return &orderService1{
		tm:          tm,
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *orderService1) CheckOut(ctx context.Context, userID int) error {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)

	if err != nil {
		return err
	}

	cartItems, err := s.cartRepo.RequestCartItems(ctx, tx, cart.CartID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return errors.New("cart empty, cant checkout")
	}

	orderItems := make([]OrderItem, len(cartItems))
	var subtotal float64

	for i, cartItem := range cartItems {
		orderItems[i] = OrderItem{
			ProductID:       cartItem.ProductID,
			Quantity:        cartItem.Quantity,
			PriceAtPurchase: cartItem.Price,
		}
		err = s.productRepo.DecreaseStock(ctx, tx, cartItem.ProductID, cartItem.Quantity)
		if err != nil {
			return err
		}
		subtotal += cartItem.Subtotal
	}

	err = s.orderRepo.CheckOut(ctx, tx, cart.UserID, subtotal, orderItems)
	if err != nil {
		return err
	}

	err = s.cartRepo.MarkCartAsCheckedOut(ctx, tx, cart.CartID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *orderService1) GetOrders(ctx context.Context, userID int) ([]OrderWithItems, error) {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	orders, err := s.orderRepo.GetOrdersByUser(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	return orders, tx.Commit()
}

func (s *orderService1) GetAllOrders(ctx context.Context) ([]UserOrders, error) {

	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	orders, err := s.orderRepo.GetAllOrders(ctx, tx)
	if err != nil {
		return nil, err
	}

	return orders, tx.Commit()
}
