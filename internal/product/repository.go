package product

import (
	"context"
	"database/sql"
	"fmt"
)

// Produkt interface
type ProductRepository interface {
	GetByProductID(id int) (Product, error)
	getAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
	GetProductStock(tx *sql.Tx, id int) (int, error)
	updateProduct(id int, stock int, price float64) error
	DecreaseStock(ctx context.Context, tx *sql.Tx, productID, quantity int) error
}

// SQL GREJS
type mysqlProductRepository struct {
	db *sql.DB
}

func NewMysqlProductRepository(db *sql.DB) *mysqlProductRepository {
	return &mysqlProductRepository{db: db}
}

// HÄMTAR PRODUKT ID
func (r *mysqlProductRepository) GetByProductID(id int) (Product, error) {
	var p Product
	err := r.db.QueryRow(`
		SELECT product_id, product_name, stock, price, manufacturer, description, category_id 
		FROM products WHERE product_id = ?`,
		id,
	).Scan(
		&p.ProductID,
		&p.ProductName,
		&p.Stock,
		&p.Price,
		&p.Manufacturer,
		&p.Description,
		&p.CategoryID,
	)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// HÄMTAR ALLA PRODUKTER
func (r *mysqlProductRepository) getAll() ([]Product, error) {
	rows, err := r.db.Query(`
		SELECT p.product_id, p.product_name, p.stock, p.price, p.manufacturer, p.description, 
		       p.category_id, c.category_name
		FROM products p
		JOIN categories c ON p.category_id = c.category_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ProductID,
			&p.ProductName,
			&p.Stock,
			&p.Price,
			&p.Manufacturer,
			&p.Description,
			&p.CategoryID,
			&p.CategoryName,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// lÄGG TILL EN NY PRODUKT
func (r *mysqlProductRepository) registerProduct(p Product) error {
	_, err := r.db.Exec(`
		INSERT INTO products (product_name, stock, price, manufacturer, description, category_id)
		VALUES (?, ?, ?, ?, ?, ?)`,
		p.ProductName,
		p.Stock,
		p.Price,
		p.Manufacturer,
		p.Description,
		p.CategoryID,
	)
	return err
}

// TA BORT PRODUKT
func (r *mysqlProductRepository) deleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE product_id = ?", id)
	return err
}

func (r *mysqlProductRepository) GetProductStock(tx *sql.Tx, id int) (int, error) {
	var stock int
	err := tx.QueryRow("SELECT stock FROM products WHERE product_id = ?", id).Scan(&stock)
	if err != nil {
		return 0, err
	}
	return stock, nil
}

// uppdatera produkters pris/stock
func (r *mysqlProductRepository) updateProduct(id int, stock int, price float64) error {
	_, err := r.db.Exec(`
		UPDATE products 
		SET stock = ?, price = ?
		WHERE product_id = ?`,
		stock, price, id,
	)
	return err
}

func (s *mysqlProductRepository) DecreaseStock(ctx context.Context, tx *sql.Tx, productID, quantity int) error {
	res, err := tx.ExecContext(
		ctx,
		`UPDATE products
		 SET stock = stock - ?
		 WHERE product_id = ? AND stock >= ?`,
		quantity,
		productID,
		quantity,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("not enough stock")
	}

	return nil
}
