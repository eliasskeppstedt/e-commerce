package product

import (
	_ "ecommerce/duckyarmy/internal/review"
	"ecommerce/duckyarmy/internal/transaction"
	"errors"
)

type productService interface {
	getByProductID(id int) (Product, error)
	getAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
	updateProduct(id int, stock int, price float64) error
}

type productServiceImp struct {
	tm   transaction.TxManager
	repo ProductRepository
}

func NewProductServiceImp(tm transaction.TxManager, r ProductRepository) *productServiceImp {
	return &productServiceImp{tm: tm, repo: r}
}

func (s *productServiceImp) getByProductID(id int) (Product, error) {
	return s.repo.GetByProductID(id)
}

func (s *productServiceImp) getAll() ([]Product, error) {
	/*products, err := s.repo.getAll()
	if err != nil {
		return nil, err
	}
	reviews := s.*/
	return s.repo.getAll()
}

// --- Add validation for negative stock/price ---
func (s *productServiceImp) registerProduct(product Product) error {
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	if product.Price < 0 {
		return errors.New("price cannot be negative")
	}
	return s.repo.registerProduct(product)
}

func (s *productServiceImp) deleteProduct(id int) error {
	return s.repo.deleteProduct(id)
}

// --- Add validation for update ---
func (s *productServiceImp) updateProduct(id int, stock int, price float64) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	return s.repo.updateProduct(id, stock, price)
}
