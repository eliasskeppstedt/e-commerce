package cart

import (
	"context"
	"ecommerce/duckyarmy/internal/product"
	"ecommerce/duckyarmy/internal/transaction"
	"errors"
)

type CartService interface {
	AddItem(ctx context.Context, userID, productID, quantity int) error
	RequestCartItems(ctx context.Context, userID int) ([]CartItemRequest, error)
	RemoveItem(ctx context.Context, userID, productID, quantity int) error
}

type cartService1 struct {
	tm          transaction.TxManager
	productRepo product.ProductRepository
	cartRepo    CartRepository
}

func NewCartService1(
	tm transaction.TxManager,
	pR product.ProductRepository,
	cR CartRepository,
) *cartService1 {
	return &cartService1{
		tm:          tm,
		productRepo: pR,
		cartRepo:    cR,
	}
}

func (s *cartService1) AddItem(ctx context.Context, userID, productID, quantity int) error {

	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)

	if err != nil {
		return err
	}

	stock, err := s.productRepo.GetProductStock(ctx, tx, productID)
	if err != nil {
		return err
	}

	curQuantity, err := s.cartRepo.GetCartItemQuantity(ctx, tx, cart.CartID, productID)
	if err != nil {
		return err
	}

	if curQuantity+quantity > stock {
		return errors.New("error: not enough in stock")
	}

	err = s.cartRepo.AddItem(ctx, tx, cart.CartID, productID, quantity)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *cartService1) RequestCartItems(ctx context.Context, userID int) ([]CartItemRequest, error) {
	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	cartItems, err := s.cartRepo.RequestCartItems(ctx, tx, cart.CartID)
	if err != nil {
		return nil, err
	}

	return cartItems, tx.Commit()
}

func (s *cartService1) RemoveItem(ctx context.Context, userID, productID, quantityToAdd int) error {

	tx, err := s.tm.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	cart, err := s.cartRepo.GetCartByUserID(ctx, tx, userID)
	if err != nil {
		return err
	}

	cartItemAmount, err := s.cartRepo.GetCartItemQuantity(ctx, tx, cart.CartID, productID)

	if err != nil {
		return err
	}

	if cartItemAmount > quantityToAdd {
		err = s.cartRepo.RemoveItem(ctx, tx, cart.CartID, productID, quantityToAdd)
	} else {
		err = s.cartRepo.DeleteItem(ctx, tx, cart.CartID, productID)
	}
	if err != nil {
		return err
	}

	return tx.Commit()
}
