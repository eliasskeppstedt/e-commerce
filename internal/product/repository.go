package product

import "database/sql"

// Produkt interface
type productRepository interface {
	getByProductID(productID int) (Product, error)
	getAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
}

// SQL GREJS
type mysqlProductRepository struct {
	db *sql.DB
}

func NewMysqlProductRepository(db *sql.DB) *mysqlProductRepository {
	return &mysqlProductRepository{db: db}
}

// HÄMTAR PRODUKT ID
func (r *mysqlProductRepository) getByProductID(productID int) (Product, error) {
	var p Product
	err := r.db.QueryRow(`
		SELECT product_id, product_name, stock, price, manufacturer, description, category_name 
		FROM products WHERE product_id = ?`,
		productID,
	).Scan(
		&p.ProductID,
		&p.ProductName,
		&p.Stock,
		&p.Price,
		&p.Manufacturer,
		&p.Description,
		&p.CategoryName,
	)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

// HÄMTAR ALLA PRODUKTER
func (r *mysqlProductRepository) getAll() ([]Product, error) {
	rows, err := r.db.Query(`
		SELECT product_id, product_name, stock, price, manufacturer, description, category_name 
		FROM products
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
		INSERT INTO products (product_name, stock, price, manufacturer, description, category_name)
		VALUES (?, ?, ?, ?, ?, ?)`,
		p.ProductName,
		p.Stock,
		p.Price,
		p.Manufacturer,
		p.Description,
		p.CategoryName,
	)
	return err
}

// TA BORT PRODUKT
func (r *mysqlProductRepository) deleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE product_id = ?", id)
	return err
}
