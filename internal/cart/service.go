package cart

import (
	"context"
	"ecommerce/duckyarmy/internal/product"
	"ecommerce/duckyarmy/internal/transaction"
	_ "errors"
)

type CartService interface {
	AddItem(ctx context.Context, carttID, productID, quantity int) error
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

	err = s.productRepo.DecreaseStock(ctx, tx, productID, quantity)
	if err != nil {
		return err
	}

	err = s.cartRepo.AddItem(ctx, tx, cart.CartID, productID, quantity)
	if err != nil {
		return err
	}

	return tx.Commit()
}
